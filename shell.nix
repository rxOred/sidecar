# shell.nix

{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.kubectl
    pkgs.k3d
    pkgs.k3s 
    pkgs.minikube
    pkgs.docker
  ];

  shellHook = ''
    # Set up Go environment variables
    export GOPATH=$PWD/go
    export PATH=$GOPATH/bin:$PATH

    echo "Kubernetes tools are ready to use."
    echo "To start Minikube, run: minikube start"
    echo "To interact with Kubernetes, use kubectl"
  '';

   # Ensure Docker daemon is running
  preShellHook = ''
    if ! systemctl is-active --quiet docker; then
      echo "Starting Docker daemon..."
      sudo systemctl start docker
    fi
  '';
}

