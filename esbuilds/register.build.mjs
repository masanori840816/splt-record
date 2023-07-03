import * as esbuild from 'esbuild';

await esbuild.build({
  entryPoints: ['ts/register.page.ts'],
  bundle: true,
  minify: false,
  outfile: 'templates/js/register.page.js',
});
