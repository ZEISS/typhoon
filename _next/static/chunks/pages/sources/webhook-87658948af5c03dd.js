(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[26],{3140:function(e,i,s){(window.__NEXT_P=window.__NEXT_P||[]).push(["/sources/webhook",function(){return s(4548)}])},4548:function(e,i,s){"use strict";s.r(i),s.d(i,{useTOC:function(){return a}});var n=s(5893),t=s(7812),h=s(5558),r=s(9299);function a(e){return[{value:"Configuration",id:"configuration",depth:2}]}i.default=(0,t.c)(function(e){let{toc:i=a(e)}=e,s={a:"a",code:"code",h2:"h2",li:"li",p:"p",pre:"pre",span:"span",ul:"ul",...(0,r.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(s.p,{children:"The source exposes an HTTP endpoint. This endpoint can be used to virtually\nintegrate with any 3rd-party system. It is a generic way to ingest events via HTTP."}),"\n",(0,n.jsx)(s.p,{children:"Every request will create an event."}),"\n",(0,n.jsxs)(s.p,{children:["The webhook can receive arbitrary data that is wrapped in a ",(0,n.jsx)(s.a,{href:"https://cloudevents.io/",children:"CloudEvent"})," envelope."]}),"\n",(0,n.jsx)(s.h2,{id:i[0].id,children:i[0].value}),"\n",(0,n.jsx)(s.pre,{tabIndex:"0","data-language":"yaml","data-word-wrap":"",children:(0,n.jsxs)(s.code,{children:[(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"apiVersion"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"sources.typhoon.zeiss.com/v1alpha1"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"kind"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"WebhookSource"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"metadata"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  name"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"sample"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"spec"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  eventType"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"com.zeiss.typhoon.sample.event"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  eventSource"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"hungry-hypatia"})]}),"\n",(0,n.jsx)(s.span,{children:" "}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  eventExtensionAttributes"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"    from"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:"    - "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"path"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:"    - "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"queries"})]}),"\n",(0,n.jsx)(s.span,{children:" "}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  basicAuthUsername"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"webhook"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  basicAuthPassword"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"    value"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"supersecret"})]}),"\n",(0,n.jsx)(s.span,{children:" "}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"  sink"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"    ref"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:":"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"      apiVersion"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"serving.knative.dev/v1"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"      kind"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"Service"})]}),"\n",(0,n.jsxs)(s.span,{children:[(0,n.jsx)(s.span,{style:{"--shiki-light":"#22863A","--shiki-dark":"#85E89D"},children:"      name"}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#24292E","--shiki-dark":"#E1E4E8"},children:": "}),(0,n.jsx)(s.span,{style:{"--shiki-light":"#032F62","--shiki-dark":"#9ECBFF"},children:"event-display"})]})]})}),"\n",(0,n.jsx)(s.p,{children:"Parameters:"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:["Name: Unique name of the ",(0,n.jsx)(s.code,{children:"WebhookSource"})," in the namespaces."]}),"\n",(0,n.jsx)(s.li,{children:"EventType: Type of the event that will be produced by the source."}),"\n",(0,n.jsx)(s.li,{children:"EventSource: Source of the event that will be produced by the source."}),"\n",(0,n.jsx)(s.li,{children:"Basic Auth Username (optional): Username for basic authentication."}),"\n",(0,n.jsx)(s.li,{children:"Basic Auth Password (optional): Password for basic authentication."}),"\n"]}),"\n",(0,n.jsx)(s.p,{children:"Events produced have the following attributes:"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:["Type of the event is defined the ",(0,n.jsx)(s.code,{children:"WebhookSource"})," configuration, e.g. ",(0,n.jsx)(s.code,{children:"com.zeiss.typhoon.sample.event"})]}),"\n",(0,n.jsxs)(s.li,{children:["Source is defined in the ",(0,n.jsx)(s.code,{children:"WebhookSource"})," configuration, e.g. ",(0,n.jsx)(s.code,{children:"hungry-hypatia"})]}),"\n",(0,n.jsxs)(s.li,{children:["Schema of the ",(0,n.jsx)(s.code,{children:"data"})," attribite depends on the client side send data to the ",(0,n.jsx)(s.code,{children:"WebhookSource"})]}),"\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.code,{children:"datacontenttype"})," is set to the ",(0,n.jsx)(s.code,{children:"Content-Type"})," received at the incoming request"]}),"\n"]})]})},"/sources/webhook",{filePath:"src/pages/sources/webhook.mdx",timestamp:1742979323e3,pageMap:h.v,frontMatter:{title:"Webhook Source (HTTP)"},title:"Webhook Source (HTTP)"},"undefined"==typeof RemoteContent?a:RemoteContent.useTOC)},5558:function(e,i,s){"use strict";s.d(i,{v:function(){return n}});let n=[{data:{index:"Home",getting_started:"Getting Started",development:"Development",sources:"Sources",targets:"Targets",transformations:"Transformations"}},{name:"development",route:"/development",children:[{data:{index:"Quickstart"}},{name:"index",route:"/development",frontMatter:{title:"Quickstart"}}]},{name:"getting_started",route:"/getting_started",children:[{data:{index:"Quickstart",concepts:"Concepts"}},{name:"concepts",route:"/getting_started/concepts",frontMatter:{title:"Concepts"}},{name:"index",route:"/getting_started",frontMatter:{title:"Quickstart"}}]},{name:"index",route:"/",frontMatter:{title:"Home"}},{name:"sources",route:"/sources",children:[{name:"index",route:"/sources",frontMatter:{title:"Quickstart"}},{name:"webhook",route:"/sources/webhook",frontMatter:{title:"Webhook Source (HTTP)"}}]},{name:"targets",route:"/targets",children:[{name:"http",route:"/targets/http",frontMatter:{title:"HTTP Target"}},{name:"index",route:"/targets",frontMatter:{title:"Quickstart"}}]},{name:"transformations",route:"/transformations",children:[{name:"index",route:"/transformations",frontMatter:{title:"Quickstart"}}]}]}},function(e){e.O(0,[7812,2888,9774,179],function(){return e(e.s=3140)}),_N_E=e.O()}]);