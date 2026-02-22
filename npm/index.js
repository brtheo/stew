const path = require('path');
const stewPath = path.dirname(require.resolve('@brtheo/stew'));
const binPath = path.resolve(stewPath, 'bin', 'stew-bin');
const { execFileSync } = require('child_process');

execFileSync(binPath, process.argv.slice(2), { stdio: 'inherit' });
