const fs = require('fs');

const [version, packagePath] = process.argv.slice(2);

const cleanVersion = version.startsWith('v') ? version.slice(1) : version;

const pkg = JSON.parse(fs.readFileSync(packagePath, 'utf8'));
pkg.version = cleanVersion;
fs.writeFileSync(packagePath, JSON.stringify(pkg, null, 2));
console.log(`Updated version to ${cleanVersion} in ${packagePath}`);