(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[7961],{4735:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/getting_started",function(){return n(6209)}])},6209:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return c},useTOC:function(){return d}});var r=n(5893),i=n(7812),a=n(5558),s=n(9299),o=n(7294);function u({chart:e}){let t=(0,o.useId)(),[i,a]=(0,o.useState)(""),s=(0,o.useRef)(null),u=function(e){let[t,n]=(0,o.useState)(!1);return(0,o.useEffect)(()=>{let t=new IntersectionObserver(([e])=>{e.isIntersecting&&(t.disconnect(),n(!0))});return t.observe(e.current),()=>{t.disconnect()}},[e]),t}(s);return(0,o.useEffect)(()=>{if(!u)return;let r=document.documentElement,i=new MutationObserver(o);return i.observe(r,{attributes:!0}),o(),()=>{i.disconnect()};async function o(){let i=r.classList.contains("dark")||r.attributes.getNamedItem("data-theme")?.value==="dark",{default:o}=await Promise.all([n.e(805),n.e(2858)]).then(n.bind(n,2858));try{o.initialize({startOnLoad:!1,securityLevel:"loose",fontFamily:"inherit",themeCSS:"margin: 1.5rem auto 0;",theme:i?"dark":"default"});let{svg:n}=await o.render(t.replaceAll(":",""),e.replaceAll("\\n","\n"),s.current);a(n)}catch(e){console.error("Error while rendering mermaid",e)}}},[e,u]),(0,r.jsx)("div",{ref:s,dangerouslySetInnerHTML:{__html:i}})}function d(e){return[{value:"Application integration",id:"application-integration",depth:2},{value:"IT Automation",id:"it-automation",depth:2},{value:"SaaS Integrations",id:"saas-integrations",depth:2}]}var c=(0,i.c)(function(e){let{toc:t=d(e)}=e,n={h1:"h1",h2:"h2",p:"p",...(0,s.a)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.h1,{children:"How it works"}),"\n",(0,r.jsx)(n.p,{children:"Typhoon is a cloud-native technology that help you to receive, filter, transform, route and deliver events\nacross a wide-range of sources and targets."}),"\n",(0,r.jsx)(u,{chart:"graph TD\\n\\nsubgraph SRC [Event Sources]\\n    A[SaaS apps]\\n    B[Microservices]\\n    C[Custom apps]\\nend\\n\\nsubgraph DEST [Targets]\\n    direction RL\\n    E[Webhook]\\n    F[API destination]\\n    G[Service targets]\\nend\\n\\nsubgraph Z [Typhoon]\\n    direction RL\\n    subgraph X [Event Bus]\\n         direction RL\\n        H[Default event bus]\\n        I[Custom event bus]\\n    end\\n    R[Rules]\\nend\\n\\nH --> R;\\nA --> H;\\nB --> H;\\nC --> H;\\nR --> E;\\nR --> F;\\nR --> G;"}),"\n",(0,r.jsx)(n.h1,{children:"Use cases"}),"\n",(0,r.jsx)(n.h2,{id:t[0].id,children:t[0].value}),"\n",(0,r.jsx)(n.p,{children:"Send events from on-premises applications to the cloud and back and use the events to build new applications."}),"\n",(0,r.jsx)(n.h2,{id:t[1].id,children:t[1].value}),"\n",(0,r.jsx)(n.p,{children:"Use events to automate your infrastructure for validating configurations, monitoring, and alerting.\nAudit user behavior, or remediate security incidents."}),"\n",(0,r.jsx)(n.h2,{id:t[2].id,children:t[2].value}),"\n",(0,r.jsx)(n.p,{children:"Receive events from your SaaS applications, like Salesforce or ServiceNow, enrich your data,\nand send the results back or to other destinations."})]})},"/getting_started",{filePath:"src/pages/getting_started/index.mdx",timestamp:1740998483e3,pageMap:a.v,frontMatter:{title:"Quickstart"},title:"Quickstart"},"undefined"==typeof RemoteContent?d:RemoteContent.useTOC)},5558:function(e,t,n){"use strict";n.d(t,{v:function(){return r}});let r=[{data:{index:"Home",getting_started:"Getting Started",development:"Development"}},{name:"development",route:"/development",children:[{data:{index:"Quickstart"}},{name:"index",route:"/development",frontMatter:{title:"Quickstart"}}]},{name:"getting_started",route:"/getting_started",children:[{data:{index:"Quickstart",concepts:"Concepts"}},{name:"concepts",route:"/getting_started/concepts",frontMatter:{title:"Concepts"}},{name:"index",route:"/getting_started",frontMatter:{title:"Quickstart"}}]},{name:"index",route:"/",frontMatter:{title:"Home"}}]}},function(e){e.O(0,[7812,2888,9774,179],function(){return e(e.s=4735)}),_N_E=e.O()}]);