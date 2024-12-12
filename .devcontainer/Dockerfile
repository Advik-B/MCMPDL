# Base image
FROM mcr.microsoft.com/devcontainers/base:dev-ubuntu-24.04

# Set non-interactive mode for apt-get
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies and add repositories
RUN apt-get update && apt-get install -y --no-install-recommends \
    software-properties-common \
    wget \
    gnupg \
    zsh && \
    # Add Kitware APT repository for the latest CMake
    wget -qO- https://apt.kitware.com/keys/kitware-archive-latest.asc | gpg --dearmor -o /usr/share/keyrings/kitware-archive-keyring.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/kitware-archive-keyring.gpg] https://apt.kitware.com/ubuntu/ jammy main" > /etc/apt/sources.list.d/kitware.list && \
    # Add LLVM APT repository for the latest Clang
    wget -qO- https://apt.llvm.org/llvm-snapshot.gpg.key | gpg --dearmor -o /usr/share/keyrings/llvm-keyring.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/llvm-keyring.gpg] http://apt.llvm.org/jammy/ llvm-toolchain-jammy main" > /etc/apt/sources.list.d/llvm.list && \
    apt-get update && \
    # Install latest versions of cmake, ninja-build, ccache, clang, git
    apt-get install -y --no-install-recommends \
    cmake \
    ninja-build \
    ccache \
    clang \
    git && \
    # Clean up
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Set zsh as the default shell
RUN chsh -s /usr/bin/zsh

# Set ccache as the default compiler cache
ENV CCACHE_DIR=/root/.ccache \
    CCACHE_BASEDIR=/workspaces

# Optional: Install Oh My Zsh for a better zsh experience
RUN wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O - | zsh || true
