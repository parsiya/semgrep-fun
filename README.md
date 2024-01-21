Fun and Weird Semgrep Experiments for the Entire Family

Blog is at: https://parsiya.net/blog/semgrep-fun/

## Setup
The experiments were done in a Debian 11 machine under WSL2. But should work on
any OS supported by Semgrep (every popular OS except Windows) and Go.

Semgrep in WSL:
https://parsiya.io/research/semgrep-tips/#semgrep-on-windows-via-wsl.

1. Clone the repository: `git clone --recurse-submodules https://github.com/parsiya/semgrep-fun`
2. Install Semgrep: `python -m pip install semgrep==1.52.0` or `pipx install semgrep==1.52.0`.
    1. Using Semgrep v1.52.0 because the structs might be different in future versions.