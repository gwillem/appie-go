test:
    go test -timeout 5s ./...

integration-test:
    go test -tags integration -count=1 -v ./...

release: test
    #!/usr/bin/env bash
    set -euo pipefail

    # Determine next version (bump patch)
    latest=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
    IFS='.' read -r major minor patch <<< "${latest#v}"
    next="v${major}.${minor}.$((patch + 1))"
    echo "Releasing ${next} (was ${latest})"

    # Build binaries
    targets=("darwin/arm64" "linux/amd64" "linux/arm64")
    artifacts=()
    for target in "${targets[@]}"; do
        os="${target%/*}"
        arch="${target#*/}"
        out="dist/appie-${os}-${arch}"
        echo "Building ${os}/${arch}..."
        GOOS="${os}" GOARCH="${arch}" go build -ldflags="-s -w" -o "${out}" ./cmd/appie
        artifacts+=("${out}")
    done

    # Tag and release
    git tag "${next}"
    git push origin "${next}"
    gh release create "${next}" "${artifacts[@]}" --generate-notes
