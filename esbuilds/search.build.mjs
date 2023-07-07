import * as esbuild from 'esbuild';

await esbuild.build({
  entryPoints: ['ts/search.page.ts'],
  bundle: true,
  minify: false,
  outfile: 'templates/js/search.page.js',
});
