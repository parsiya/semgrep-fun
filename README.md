Fun and Weird Semgrep Experiments for the Entire Family

ZZZ import `semgrep_go` as a module later.

## Setup
The experiments were done in a Debian 11 machine under WSL2. But should work on
any OS supported by Semgrep (every popular OS except Windows) and Go.

Semgrep in WSL:
https://parsiya.io/research/semgrep-tips/#semgrep-on-windows-via-wsl.

ZZ add note about the latest Semgrep version that I try here. The output
structure might be modified in future versions and not usable.

1. Clone the repository: `git clone --recurse-submodules https://github.com/parsiya/ZZZ`
2. Install Semgrep: `python -m pip install semgrep`.
    1. Or update: `python -m pip install Semgrep -U`.
3. Get submodules: `git submodule update --init`