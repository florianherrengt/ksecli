# KSECLI

Enforce the list of secrets in Kubernetes.

This is the equivalent of doing:
- `echo -n 'secret' | base64 | pbcopy`
- Pasting to the corresponding value in secrets.yml

The script will only ask for missing values. To edit one, remove it from `secrets.yml`.