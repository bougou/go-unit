#!/usr/bin/env node
/**
 * Verify internal links in the built docs site (no 404s).
 * Run after `pnpm build`.
 */

import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const distDir = path.join(__dirname, '..', 'dist');
const basePath = '/go-unit';

const ASSET_EXTENSIONS = new Set([
  '.css',
  '.js',
  '.mjs',
  '.xml',
  '.json',
  '.png',
  '.svg',
  '.ico',
  '.webp',
  '.woff',
  '.woff2',
  '.webmanifest',
]);

const errors = [];
const checked = new Set();

function walk(dir, files = []) {
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) walk(full, files);
    else if (entry.name.endsWith('.html')) files.push(full);
  }
  return files;
}

function pageUrlFromFile(file) {
  const rel = path.relative(distDir, file);
  if (rel === 'index.html') return `${basePath}/`;
  if (rel.endsWith('/index.html')) {
    return `${basePath}/${rel.slice(0, -'index.html'.length)}`;
  }
  return `${basePath}/${rel}`;
}

function isSkippableHref(href) {
  if (!href || href.startsWith('#')) return true;
  if (href.startsWith('mailto:') || href.startsWith('tel:') || href.startsWith('javascript:'))
    return true;
  if (href.startsWith('http://') || href.startsWith('https://')) return true;

  const pathname = href.split('#')[0].split('?')[0];
  const ext = path.extname(pathname);
  return ASSET_EXTENSIONS.has(ext);
}

function resolveHref(pageUrl, href) {
  const base = `http://localhost${pageUrl.endsWith('/') ? pageUrl : `${pageUrl}/`}`;
  const resolved = new URL(href, base);
  let pathname = resolved.pathname;
  if (!pathname.endsWith('/')) pathname += '/';
  return { pathname, hash: resolved.hash ? resolved.hash.slice(1) : null };
}

function resolveToFile(pathname) {
  let p = pathname;
  if (p.startsWith(basePath)) p = p.slice(basePath.length);
  const rel = p.replace(/^\//, '');
  if (!rel) return path.join(distDir, 'index.html');

  const asIndex = path.join(distDir, rel, 'index.html');
  if (fs.existsSync(asIndex)) return asIndex;

  const asFile = path.join(distDir, rel);
  if (fs.existsSync(asFile) && fs.statSync(asFile).isFile()) return asFile;

  return null;
}

function hasAnchor(filePath, anchor) {
  const html = fs.readFileSync(filePath, 'utf8');
  const id = anchor.toLowerCase();
  return [
    new RegExp(`\\bid=["']${escapeRegExp(id)}["']`, 'i'),
    new RegExp(`\\bid=["']user-content-${escapeRegExp(id)}["']`, 'i'),
    new RegExp(`\\bname=["']${escapeRegExp(id)}["']`, 'i'),
  ].some((re) => re.test(html));
}

function escapeRegExp(value) {
  return value.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

function checkLink(pageUrl, href) {
  const key = `${pageUrl} -> ${href}`;
  if (checked.has(key)) return;
  checked.add(key);

  const { pathname, hash } = resolveHref(pageUrl, href);
  const file = resolveToFile(pathname);
  if (!file) {
    errors.push({ page: pageUrl, href, message: `missing page ${pathname}` });
    return;
  }
  if (hash && !hasAnchor(file, hash)) {
    errors.push({
      page: pageUrl,
      href,
      message: `anchor #${hash} not found in ${path.relative(distDir, file)}`,
    });
  }
}

if (!fs.existsSync(distDir)) {
  console.error('dist/ not found — run pnpm build first');
  process.exit(1);
}

const htmlFiles = walk(distDir);
const hrefPattern = /\bhref=["']([^"']+)["']/gi;

for (const file of htmlFiles) {
  const pageUrl = pageUrlFromFile(file);
  const html = fs.readFileSync(file, 'utf8');
  let match;
  while ((match = hrefPattern.exec(html))) {
    const href = match[1];
    if (isSkippableHref(href)) continue;
    checkLink(pageUrl, href);
  }
}

// Landing pages (static checks)
for (const href of [
  `${basePath}/en/concepts/dimensions/`,
  `${basePath}/en/guides/typed-quantities/`,
  `${basePath}/en/concepts/`,
  `${basePath}/zh/concepts/`,
  `${basePath}/zh/concepts/dimensions/`,
  `${basePath}/zh/guides/typed-quantities/`,
  `${basePath}/`,
]) {
  checkLink('landing', href);
}

if (errors.length === 0) {
  console.log(`✓ All internal links OK (${htmlFiles.length} HTML pages checked)`);
  process.exit(0);
}

console.error(`✗ ${errors.length} broken internal link(s):\n`);
for (const error of errors) {
  console.error(`  ${error.page}`);
  console.error(`    ${error.href}`);
  console.error(`    → ${error.message}\n`);
}
process.exit(1);
