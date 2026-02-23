// Just install for darwin for sake of simplicity, sorry.
// Naively include both am64/arm64 arch in the node package.

// maps process.arch to GOARCH
let GOARCH_MAP = {
  'arm64': 'arm64',
  'x64': 'amd64'
};

let GOOS_MAP = {
  'darwin': 'darwin',
  'win32': 'windows',
  'win64': 'windows',
  'linux': 'linux'
};

if (!(process.arch in GOARCH_MAP)) {
  console.error(`Sorry this is only packaged for ${Object.keys(GOARCH_MAP)} at the moment. Current arch: ${process.arch}`);
  process.exit(1);
}

if (!(process.platform in GOOS_MAP)) {
  console.error(`Sorry this is only packaged for ${Object.keys(GOOS_MAP)} at the moment. Current os : ${process.platform}`);
  process.exit(1);
}

const arch = GOARCH_MAP[process.arch];
const platform = GOOS_MAP[process.platform];
const ext = process.platform == 'windows' ? '.exe' : ''
const installTarget = `stew-${platform}-${arch}${ext}`;

// "Install"
const { exec } = require('child_process');
exec(`${platform == 'windows' ? 'copy' : 'cp'} bin/${installTarget} bin/stew-bin`, (err) => {
  if (err) {
    console.error(err);
    process.exit(1);
  }
});
