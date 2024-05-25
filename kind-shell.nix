with import <nixpkgs> {};

mkShell {
  buildInputs = [
    go
    kubectl
    docker
    docker-compose
    conntrack-tools
    curl
  ];

  shellHook = ''
    export GOPATH=$PWD/go
    export PATH=$GOPATH/bin:$PATH

    # Ensure Docker daemon is running
    sudo systemctl start docker || true

    # Ensure user is in the Docker group
    if ! groups | grep -q docker; then
      sudo usermod -aG docker $USER
      echo "You need to log out and log back in for Docker group changes to take effect."
    fi

    # Ensure Docker is running
    sudo systemctl is-active --quiet docker || sudo systemctl start docker

    # Install Kind if not already installed
    if ! command -v kind &> /dev/null; then
      curl -Lo ./bins/kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64 
      chmod +x ./bins/kind
      export PATH=$PWD/bins:$PATH
    fi
  '';
}

