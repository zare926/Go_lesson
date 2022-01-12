# npm ERR! ERESOLVE unable to resolve dependency treeの対処方法
SPAのブラウザバックのscroll復活をしたく、ためにしにreact-scroll-restorationを導入しようとしたらタイトルのエラーがでました。<br>
[react-scroll-restoration](https://www.npmjs.com/package/react-scroll-restoration)
## 結論
`-save --legacy-peer-deps`をnpm installのあとにつければいいそうです。<br>

## 実際のエラー
```
$ npm install react-scroll-restoration
npm ERR! code ERESOLVE
npm ERR! ERESOLVE unable to resolve dependency tree
npm ERR! 
npm ERR! While resolving: testApp@0.1.0
npm ERR! Found: react-router-dom@5.3.0
npm ERR! node_modules/react-router-dom
npm ERR!   react-router-dom@"^5.2.0" from the root project
npm ERR! 
npm ERR! Could not resolve dependency:
npm ERR! peer react-router-dom@"^4.0.0" from react-scroll-restoration@1.0.8
npm ERR! node_modules/react-scroll-restoration
npm ERR!   react-scroll-restoration@"*" from the root project
npm ERR! 
npm ERR! Fix the upstream dependency conflict, or retry
npm ERR! this command with --force, or --legacy-peer-deps
npm ERR! to accept an incorrect (and potentially broken) dependency resolution.
npm ERR! 
npm ERR! See /Users/userName/.npm/eresolve-report.txt for a full report.

npm ERR! A complete log of this run can be found in:
npm ERR!     /Users/userName/.npm/_logs/2021-12-28T11_47_37_669Z-debug.log
```
npm ERR! Found: react-router-dom@5.3.0<br>
↑因果関係があるライブラリの今使ってるバージョン<br>
npm ERR! peer react-router-dom@"^4.0.0" from react-scroll-restoration@1.0.8<br>
↑実際にさぽーとサポートしているバージョン<br>
```
$ npm install react-scroll-restoration -save --legacy-peer-deps

npm WARN EBADENGINE Unsupported engine {
npm WARN EBADENGINE   package: '@jest/types@27.0.6',
npm WARN EBADENGINE   required: { node: '^10.13.0 || ^12.13.0 || ^14.15.0 || >=15.0.0' },
npm WARN EBADENGINE   current: { node: 'v12.10.0', npm: '7.16.0' }
npm WARN EBADENGINE }
npm WARN EBADENGINE Unsupported engine {
npm WARN EBADENGINE   package: 'diff-sequences@27.0.6',
npm WARN EBADENGINE   required: { node: '^10.13.0 || ^12.13.0 || ^14.15.0 || >=15.0.0' },
npm WARN EBADENGINE   current: { node: 'v12.10.0', npm: '7.16.0' }
npm WARN EBADENGINE }
npm WARN EBADENGINE Unsupported engine {
npm WARN EBADENGINE   package: 'jest-diff@27.0.6',
npm WARN EBADENGINE   required: { node: '^10.13.0 || ^12.13.0 || ^14.15.0 || >=15.0.0' },
npm WARN EBADENGINE   current: { node: 'v12.10.0', npm: '7.16.0' }
npm WARN EBADENGINE }
npm WARN EBADENGINE Unsupported engine {
npm WARN EBADENGINE   package: 'jest-get-type@27.0.6',
npm WARN EBADENGINE   required: { node: '^10.13.0 || ^12.13.0 || ^14.15.0 || >=15.0.0' },
npm WARN EBADENGINE   current: { node: 'v12.10.0', npm: '7.16.0' }
npm WARN EBADENGINE }
npm WARN EBADENGINE Unsupported engine {
npm WARN EBADENGINE   package: 'pretty-format@27.0.6',
npm WARN EBADENGINE   required: { node: '^10.13.0 || ^12.13.0 || ^14.15.0 || >=15.0.0' },
npm WARN EBADENGINE   current: { node: 'v12.10.0', npm: '7.16.0' }
npm WARN EBADENGINE }

added 1 package, and audited 2569 packages in 13s

166 packages are looking for funding
  run `npm fund` for details

43 vulnerabilities (27 moderate, 13 high, 3 critical)

To address issues that do not require attention, run:
  npm audit fix

To address all issues (including breaking changes), run:
  npm audit fix --force

Run `npm audit` for details.
```
無理やり入れているのでWARN出まくってますが、とりあえずinstallはできました。<br>
余談ですが、インストールしようとしていたreact-scroll-restorationはうまく機能しませんでした。<br>