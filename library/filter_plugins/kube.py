from ansible.errors import AnsibleFilterTypeError
import json

def cmdline(args):
  if not isinstance(args, dict):
    raise AnsibleFilterTypeError('must be a dict')
  return " ".join([
    f"--{key}={json.dumps(value)}"
    for key, value in args.items()
  ])

def etcd_peer(name):
  return f"{name}=https://{name}:2380"

class FilterModule(object):
  def filters(self):
    return dict(
      cmdline=cmdline,
      etcd_peer=etcd_peer,
    )
