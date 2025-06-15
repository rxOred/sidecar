CLUSTER_NAME="tekton-demo"
NAMESPACE="tekton-pipelines"

echo "🔧 Creating Kind cluster: $CLUSTER_NAME..."
kind create cluster --name "$CLUSTER_NAME" --image kindest/node:v1.28.0

echo "📦 Installing Tekton Pipelines..."
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

echo "⌛ Waiting for Tekton pods in '$NAMESPACE' to appear..."
sleep 10

echo "⏳ Waiting for Tekton pods to become ready..."
kubectl wait pod --all --for=condition=Ready -n "$NAMESPACE" --timeout=180s || true

echo "✅ Tekton is ready and running on Kind cluster: $CLUSTER_NAME"
