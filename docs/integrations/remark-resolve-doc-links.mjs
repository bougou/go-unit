import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { visit } from 'unist-util-visit';

const contentRoot = fileURLToPath(new URL('../src/content/docs', import.meta.url));

/**
 * Resolve Markdown links relative to the source file and prefix absolute /docs paths with base.
 */
export function remarkResolveDocLinks({ base = '' } = {}) {
  const basePrefix = base.endsWith('/') ? base.slice(0, -1) : base;

  return (tree, file) => {
    const filePath = file.history?.[0];
    if (!filePath?.includes(`${path.sep}content${path.sep}docs${path.sep}`)) return;

    visit(tree, 'link', (node) => {
      const href = node.url;
      if (
        !href ||
        href.startsWith('http://') ||
        href.startsWith('https://') ||
        href.startsWith('mailto:') ||
        href.startsWith('tel:')
      ) {
        return;
      }

      const hashIndex = href.indexOf('#');
      const linkPath = hashIndex === -1 ? href : href.slice(0, hashIndex);
      const hash = hashIndex === -1 ? '' : href.slice(hashIndex);

      if (!linkPath) {
        return;
      }

      if (linkPath.startsWith('/')) {
        if (basePrefix && !linkPath.startsWith(basePrefix)) {
          node.url = basePrefix + linkPath + hash;
        }
        return;
      }

      const sourceDir = path.dirname(filePath);
      let target = path.normalize(path.join(sourceDir, linkPath));

      if (!target.endsWith('.md')) {
        target = target.replace(/\/$/, '') + '.md';
      }

      if (!target.startsWith(contentRoot)) {
        return;
      }

      let rel = path.relative(contentRoot, target.slice(0, -'.md'.length));
      if (rel.endsWith(`${path.sep}index`) || rel === 'index') {
        rel = rel.replace(new RegExp(`${path.sep}index$`), '').replace(/^index$/, '');
      }

      let urlPath = `/${rel.replace(/\\/g, '/')}`;
      if (!urlPath.endsWith('/')) {
        urlPath += '/';
      }

      node.url = `${basePrefix}${urlPath}${hash}`;
    });
  };
}
