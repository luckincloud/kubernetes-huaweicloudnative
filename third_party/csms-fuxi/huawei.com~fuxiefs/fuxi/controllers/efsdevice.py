# -*- encoding:utf-8 -*-

import os
import fcntl
import psutil
import commands
from common.config import CONF
from oslo_concurrency import processutils
from oslo_log import log as logging
from oslo_utils import excutils
from common.i18n import _LI, _LE

LOG = logging.getLogger(__name__)


class Partition(object):
    def __init__(self, device, mountpoint, fstype, opts):
        self.device = device
        self.mountpoint = mountpoint
        self.fstype = fstype
        self.opts = opts

    def __repr__(self, *args, **kwargs):
        return str(self.__dict__)


class EfsDeviceManager(object):
    def mount_efs(self, dev_path, mount_point, share_proto):
        try:
            if share_proto == 'SMB' or share_proto == "CIFS":
                processutils.execute('mount', '-t', 'cifs', '-o', 'nocase', '-o', 'username=*', '-o',
                                     'password=*', dev_path, mount_point, run_as_root=True)
            else:
                processutils.execute('mount', '-t', 'nfs', '-o', 'vers=3', '-o', 'nolock', '-o', 'noresvport', dev_path,
                                     mount_point, run_as_root=True)
            msg = _LI("Succeed to mount the {0} to {1}").format(dev_path, mount_point)
            LOG.info(msg)
        except processutils.ProcessExecutionError as e:
            LOG.exception(e.message)
            with excutils.save_and_reraise_exception():
                msg = _LE("Unexpected error while mount sfs device. "
                          "dev_path: {0}, "
                          "mount_point: {1} "
                          "error: {2}").format(dev_path,
                                               mount_point,
                                               e)
                LOG.error(msg)

    def unmount(self, mount_point):
        if CONF.provider_type == 'storage_manager':
            lock_file = CONF.storage_manager.lock_path + '.unmount'
            try:
                if not os.path.exists(lock_file):
                    msg = _LE("Unmount lock is not exists!")
                    LOG.error(msg)
                    return
            except Exception as ex:
                msg = _LE("Unexpected err while check unmount lock. "
                          "Error: {0}").format(ex)
                LOG.error(msg)
                return

            with open(lock_file, 'w') as lockfile:
                fcntl.flock(lockfile, fcntl.LOCK_EX)
                try:
                    command = "umount -l " + mount_point
                    status, output = EfsDeviceManager().exec_shell_cmd(command, 5)
                    if status == 0:
                        return
                    else:
                        LOG.warn(output)
                        return

                except processutils.ProcessExecutionError as e:
                    LOG.exception(e.message)
                    with excutils.save_and_reraise_exception():
                        msg = _LE("Unexpected err while unmount block device. "
                                  "Mount Point: {0}, "
                                  "Error: {1}").format(mount_point, e)
                        LOG.error(msg)
                finally:
                    fcntl.flock(lockfile, fcntl.LOCK_UN)
        else:
            try:
                processutils.execute('umount', '-l', mount_point, run_as_root=True)
            except Exception as e:
                with excutils.save_and_reraise_exception():
                    msg = _LE("Unexpected err while unmount block device. "
                              "Mount Point: {0}, "
                              "Error: {1}").format(mount_point, e)
                    LOG.error(msg)

    def exec_shell_cmd(self, command, timeout=0):
        target_cmd = command
        if timeout > 0:
            target_cmd = "timeout %d %s" % (timeout, command)
        status, outputs = commands.getstatusoutput(target_cmd)
        if status == 0:
            return status, None
        return status, outputs

    def get_mounts(self):
        mounts = psutil.disk_partitions()
        return [Partition(mount.device, mount.mountpoint,
                          mount.fstype, mount.opts) for mount in mounts]


def _check_already_mount(devpath, mountpoint):
    partitions = EfsDeviceManager().get_mounts()
    for p in partitions:
        if devpath == p.device and mountpoint == p.mountpoint:
            return True
    return False


def do_mount(devpath, mountpoint, share_proto):
    if not _check_already_mount(devpath, mountpoint):
        ndm = EfsDeviceManager()
        try:
            ndm.mount_efs(devpath, mountpoint, share_proto)
            msg = _LI("mount sfs device success")
            return mountpoint, msg
        except processutils.ProcessExecutionError as e:
            LOG.exception(e.message)
            return None, _LE("Failed to mount  sfs device,mount command Failed")
    return mountpoint, _LI("the sfs device have been mounted ")


def do_unmount(mount_point):
    EfsDeviceManager().unmount(mount_point)
