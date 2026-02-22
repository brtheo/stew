const path = require('path');
const binPath = path.resolve('bin', 'stew-bin');
const { execFileSync } = require('child_process');

execFileSync(binPath, process.argv.slice(2), { stdio: 'inherit' });
