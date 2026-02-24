const fs = require('fs');
const path = require('path');

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
  console.error(`Sorry this is only packaged for ${Object.keys(GOARCH_MAP).join(', ')} at the moment. Current arch: ${process.arch}`);
  process.exit(1);
}

if (!(process.platform in GOOS_MAP)) {
  console.error(`Sorry this is only packaged for ${Object.keys(GOOS_MAP).join(', ')} at the moment. Current os : ${process.platform}`);
  process.exit(1);
}

const arch = GOARCH_MAP[process.arch];
const platform = GOOS_MAP[process.platform];
const ext = platform == 'windows' ? '.exe' : ''
const installTarget = `stew-${platform}-${arch}${ext}`;

const sourcePath = path.join(__dirname, 'bin', installTarget);
const destPath = path.join(__dirname, 'bin', `stew-bin${ext}`);

fs.copyFileSync(sourcePath, destPath);
