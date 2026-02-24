#!/usr/bin/env node
const path = require('path');

const GOOS_MAP = {
  'darwin': 'darwin',
  'win32': 'windows',
  'win64': 'windows',
  'linux': 'linux'
};
const platform = GOOS_MAP[process.platform];
const ext = platform == 'windows' ? '.exe' : ''

const stewPath = path.dirname(require.resolve('@brtheo/stew'));
const binPath = path.resolve(stewPath, 'bin', `stew-bin${ext}`);
const { execFileSync } = require('child_process');

execFileSync(binPath, process.argv.slice(2), { stdio: 'inherit' });
