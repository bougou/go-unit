import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import starlightThemeNext from 'starlight-theme-next';
import { remarkResolveDocLinks } from './integrations/remark-resolve-doc-links.mjs';

const base = '/go-unit/';

export default defineConfig({
  site: 'https://bougou.github.io',
  base,
  trailingSlash: 'always',
  markdown: {
    remarkPlugins: [[remarkResolveDocLinks, { base }]],
  },
  integrations: [
    starlight({
      title: 'go-unit',
      description:
        'Physical quantities with SI-aware units, dimensions, and conversions for Go',
      defaultLocale: 'en',
      locales: {
        en: {
          label: 'English',
          lang: 'en',
        },
        zh: {
          label: '中文',
          lang: 'zh-CN',
        },
      },
      social: [
        {
          icon: 'github',
          label: 'GitHub',
          href: 'https://github.com/bougou/go-unit',
        },
      ],
      editLink: {
        baseUrl: 'https://github.com/bougou/go-unit/edit/units/docs/',
      },
      customCss: ['./src/styles/custom.css'],
      plugins: [starlightThemeNext()],
      sidebar: [
        {
          label: 'Concepts',
          translations: { 'zh-CN': '概念' },
          autogenerate: { directory: 'concepts' },
        },
        {
          label: 'Guides',
          translations: { 'zh-CN': '指南' },
          autogenerate: { directory: 'guides' },
        },
        {
          label: 'Reference',
          translations: { 'zh-CN': '参考' },
          autogenerate: { directory: 'reference' },
        },
      ],
    }),
  ],
});
