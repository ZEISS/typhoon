(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[7812],{4080:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),Object.defineProperty(t,"Image",{enumerable:!0,get:function(){return E}});let n=i(8754),o=i(1757),r=i(5893),s=o._(i(7294)),a=n._(i(3935)),l=n._(i(3867)),h=i(5283),d=i(6594),c=i(3945);i(3179);let u=i(1928),m=n._(i(3872)),g={deviceSizes:[640,750,828,1080,1200,1920,2048,3840],imageSizes:[16,32,48,64,96,128,256,384],path:"/typhoon/_next/image",loader:"default",dangerouslyAllowSVG:!1,unoptimized:!0};function f(e,t,i,n,o,r,s){let a=null==e?void 0:e.src;e&&e["data-loaded-src"]!==a&&(e["data-loaded-src"]=a,("decode"in e?e.decode():Promise.resolve()).catch(()=>{}).then(()=>{if(e.parentElement&&e.isConnected){if("empty"!==t&&o(!0),null==i?void 0:i.current){let t=new Event("load");Object.defineProperty(t,"target",{writable:!1,value:e});let n=!1,o=!1;i.current({...t,nativeEvent:t,currentTarget:e,target:e,isDefaultPrevented:()=>n,isPropagationStopped:()=>o,persist:()=>{},preventDefault:()=>{n=!0,t.preventDefault()},stopPropagation:()=>{o=!0,t.stopPropagation()}})}(null==n?void 0:n.current)&&n.current(e)}}))}function p(e){return s.use?{fetchPriority:e}:{fetchpriority:e}}let v=(0,s.forwardRef)((e,t)=>{let{src:i,srcSet:n,sizes:o,height:a,width:l,decoding:h,className:d,style:c,fetchPriority:u,placeholder:m,loading:g,unoptimized:v,fill:w,onLoadRef:E,onLoadingCompleteRef:b,setBlurComplete:y,setShowAltText:S,sizesInput:L,onLoad:z,onError:C,...I}=e;return(0,r.jsx)("img",{...I,...p(u),loading:g,width:l,height:a,decoding:h,"data-nimg":w?"fill":"1",className:d,style:c,sizes:o,srcSet:n,src:i,ref:(0,s.useCallback)(e=>{t&&("function"==typeof t?t(e):"object"==typeof t&&(t.current=e)),e&&(C&&(e.src=e.src),e.complete&&f(e,m,E,b,y,v,L))},[i,m,E,b,y,C,v,L,t]),onLoad:e=>{f(e.currentTarget,m,E,b,y,v,L)},onError:e=>{S(!0),"empty"!==m&&y(!0),C&&C(e)}})});function w(e){let{isAppRouter:t,imgAttributes:i}=e,n={as:"image",imageSrcSet:i.srcSet,imageSizes:i.sizes,crossOrigin:i.crossOrigin,referrerPolicy:i.referrerPolicy,...p(i.fetchPriority)};return t&&a.default.preload?(a.default.preload(i.src,n),null):(0,r.jsx)(l.default,{children:(0,r.jsx)("link",{rel:"preload",href:i.srcSet?void 0:i.src,...n},"__nimg-"+i.src+i.srcSet+i.sizes)})}let E=(0,s.forwardRef)((e,t)=>{let i=(0,s.useContext)(u.RouterContext),n=(0,s.useContext)(c.ImageConfigContext),o=(0,s.useMemo)(()=>{var e;let t=g||n||d.imageConfigDefault,i=[...t.deviceSizes,...t.imageSizes].sort((e,t)=>e-t),o=t.deviceSizes.sort((e,t)=>e-t),r=null==(e=t.qualities)?void 0:e.sort((e,t)=>e-t);return{...t,allSizes:i,deviceSizes:o,qualities:r}},[n]),{onLoad:a,onLoadingComplete:l}=e,f=(0,s.useRef)(a);(0,s.useEffect)(()=>{f.current=a},[a]);let p=(0,s.useRef)(l);(0,s.useEffect)(()=>{p.current=l},[l]);let[E,b]=(0,s.useState)(!1),[y,S]=(0,s.useState)(!1),{props:L,meta:z}=(0,h.getImgProps)(e,{defaultLoader:m.default,imgConf:o,blurComplete:E,showAltText:y});return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(v,{...L,unoptimized:z.unoptimized,placeholder:z.placeholder,fill:z.fill,onLoadRef:f,onLoadingCompleteRef:p,setBlurComplete:b,setShowAltText:S,sizesInput:e.sizes,ref:t}),z.priority?(0,r.jsx)(w,{isAppRouter:!i,imgAttributes:L}):null]})});("function"==typeof t.default||"object"==typeof t.default&&null!==t.default)&&void 0===t.default.__esModule&&(Object.defineProperty(t.default,"__esModule",{value:!0}),Object.assign(t.default,t),e.exports=t.default)},5283:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),Object.defineProperty(t,"getImgProps",{enumerable:!0,get:function(){return a}}),i(3179);let n=i(6630),o=i(6594);function r(e){return void 0!==e.default}function s(e){return void 0===e?e:"number"==typeof e?Number.isFinite(e)?e:NaN:"string"==typeof e&&/^[0-9]+$/.test(e)?parseInt(e,10):NaN}function a(e,t){var i,a;let l,h,d,{src:c,sizes:u,unoptimized:m=!1,priority:g=!1,loading:f,className:p,quality:v,width:w,height:E,fill:b=!1,style:y,overrideSrc:S,onLoad:L,onLoadingComplete:z,placeholder:C="empty",blurDataURL:I,fetchPriority:M,decoding:D="async",layout:x,objectFit:O,objectPosition:N,lazyBoundary:A,lazyRoot:j,...T}=e,{imgConf:P,showAltText:k,blurComplete:R,defaultLoader:U}=t,_=P||o.imageConfigDefault;if("allSizes"in _)l=_;else{let e=[..._.deviceSizes,..._.imageSizes].sort((e,t)=>e-t),t=_.deviceSizes.sort((e,t)=>e-t),n=null==(i=_.qualities)?void 0:i.sort((e,t)=>e-t);l={..._,allSizes:e,deviceSizes:t,qualities:n}}if(void 0===U)throw Error("images.loaderFile detected but the file is missing default export.\nRead more: https://nextjs.org/docs/messages/invalid-images-config");let W=T.loader||U;delete T.loader,delete T.srcSet;let $="__next_img_default"in W;if($){if("custom"===l.loader)throw Error('Image with src "'+c+'" is missing "loader" prop.\nRead more: https://nextjs.org/docs/messages/next-image-missing-loader')}else{let e=W;W=t=>{let{config:i,...n}=t;return e(n)}}if(x){"fill"===x&&(b=!0);let e={intrinsic:{maxWidth:"100%",height:"auto"},responsive:{width:"100%",height:"auto"}}[x];e&&(y={...y,...e});let t={responsive:"100vw",fill:"100vw"}[x];t&&!u&&(u=t)}let F="",G=s(w),B=s(E);if("object"==typeof(a=c)&&(r(a)||void 0!==a.src)){let e=r(c)?c.default:c;if(!e.src)throw Error("An object should only be passed to the image component src parameter if it comes from a static image import. It must include src. Received "+JSON.stringify(e));if(!e.height||!e.width)throw Error("An object should only be passed to the image component src parameter if it comes from a static image import. It must include height and width. Received "+JSON.stringify(e));if(h=e.blurWidth,d=e.blurHeight,I=I||e.blurDataURL,F=e.src,!b){if(G||B){if(G&&!B){let t=G/e.width;B=Math.round(e.height*t)}else if(!G&&B){let t=B/e.height;G=Math.round(e.width*t)}}else G=e.width,B=e.height}}let H=!g&&("lazy"===f||void 0===f);(!(c="string"==typeof c?c:F)||c.startsWith("data:")||c.startsWith("blob:"))&&(m=!0,H=!1),l.unoptimized&&(m=!0),$&&c.endsWith(".svg")&&!l.dangerouslyAllowSVG&&(m=!0),g&&(M="high");let Z=s(v),Y=Object.assign(b?{position:"absolute",height:"100%",width:"100%",left:0,top:0,right:0,bottom:0,objectFit:O,objectPosition:N}:{},k?{}:{color:"transparent"},y),q=R||"empty"===C?null:"blur"===C?'url("data:image/svg+xml;charset=utf-8,'+(0,n.getImageBlurSvg)({widthInt:G,heightInt:B,blurWidth:h,blurHeight:d,blurDataURL:I||"",objectFit:Y.objectFit})+'")':'url("'+C+'")',V=q?{backgroundSize:Y.objectFit||"cover",backgroundPosition:Y.objectPosition||"50% 50%",backgroundRepeat:"no-repeat",backgroundImage:q}:{},K=function(e){let{config:t,src:i,unoptimized:n,width:o,quality:r,sizes:s,loader:a}=e;if(n)return{src:i,srcSet:void 0,sizes:void 0};let{widths:l,kind:h}=function(e,t,i){let{deviceSizes:n,allSizes:o}=e;if(i){let e=/(^|\s)(1?\d?\d)vw/g,t=[];for(let n;n=e.exec(i);n)t.push(parseInt(n[2]));if(t.length){let e=.01*Math.min(...t);return{widths:o.filter(t=>t>=n[0]*e),kind:"w"}}return{widths:o,kind:"w"}}return"number"!=typeof t?{widths:n,kind:"w"}:{widths:[...new Set([t,2*t].map(e=>o.find(t=>t>=e)||o[o.length-1]))],kind:"x"}}(t,o,s),d=l.length-1;return{sizes:s||"w"!==h?s:"100vw",srcSet:l.map((e,n)=>a({config:t,src:i,quality:r,width:e})+" "+("w"===h?e:n+1)+h).join(", "),src:a({config:t,src:i,quality:r,width:l[d]})}}({config:l,src:c,unoptimized:m,width:G,quality:Z,sizes:u,loader:W});return{props:{...T,loading:H?"lazy":f,fetchPriority:M,width:G,height:B,decoding:D,className:p,style:{...Y,...V},sizes:K.sizes,srcSet:K.srcSet,src:S||K.src},meta:{unoptimized:m,priority:g,placeholder:C,fill:b}}}},6630:function(e,t){"use strict";function i(e){let{widthInt:t,heightInt:i,blurWidth:n,blurHeight:o,blurDataURL:r,objectFit:s}=e,a=n?40*n:t,l=o?40*o:i,h=a&&l?"viewBox='0 0 "+a+" "+l+"'":"";return"%3Csvg xmlns='http://www.w3.org/2000/svg' "+h+"%3E%3Cfilter id='b' color-interpolation-filters='sRGB'%3E%3CfeGaussianBlur stdDeviation='20'/%3E%3CfeColorMatrix values='1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0 0 100 -1' result='s'/%3E%3CfeFlood x='0' y='0' width='100%25' height='100%25'/%3E%3CfeComposite operator='out' in='s'/%3E%3CfeComposite in2='SourceGraphic'/%3E%3CfeGaussianBlur stdDeviation='20'/%3E%3C/filter%3E%3Cimage width='100%25' height='100%25' x='0' y='0' preserveAspectRatio='"+(h?"none":"contain"===s?"xMidYMid":"cover"===s?"xMidYMid slice":"none")+"' style='filter: url(%23b);' href='"+r+"'/%3E%3C/svg%3E"}Object.defineProperty(t,"__esModule",{value:!0}),Object.defineProperty(t,"getImageBlurSvg",{enumerable:!0,get:function(){return i}})},6210:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),function(e,t){for(var i in t)Object.defineProperty(e,i,{enumerable:!0,get:t[i]})}(t,{default:function(){return l},getImageProps:function(){return a}});let n=i(8754),o=i(5283),r=i(4080),s=n._(i(3872));function a(e){let{props:t}=(0,o.getImgProps)(e,{defaultLoader:s.default,imgConf:{deviceSizes:[640,750,828,1080,1200,1920,2048,3840],imageSizes:[16,32,48,64,96,128,256,384],path:"/typhoon/_next/image",loader:"default",dangerouslyAllowSVG:!1,unoptimized:!0}});for(let[e,i]of Object.entries(t))void 0===i&&delete t[e];return{props:t}}let l=r.Image},3872:function(e,t){"use strict";function i(e){var t;let{config:i,src:n,width:o,quality:r}=e,s=r||(null==(t=i.qualities)?void 0:t.reduce((e,t)=>Math.abs(t-75)<Math.abs(e-75)?t:e))||75;return i.path+"?url="+encodeURIComponent(n)+"&w="+o+"&q="+s}Object.defineProperty(t,"__esModule",{value:!0}),Object.defineProperty(t,"default",{enumerable:!0,get:function(){return n}}),i.__next_img_default=!0;let n=i},5675:function(e,t,i){e.exports=i(6210)},9299:function(e,t,i){"use strict";i.d(t,{a:function(){return P}});var n=i(1151),o=i(5893),r=i(7294),s=i(3935);let a=(e,t)=>e===t?.tagName?.toUpperCase?.(),l=e=>a("DIV",e)||a("SPAN",e),h=e=>a("IMG",e),d=e=>e.complete&&0!==e.naturalHeight,c=e=>a("SVG",e),u=({height:e,offset:t,width:i})=>Math.min((window.innerWidth-2*t)/i,(window.innerHeight-2*t)/e),m=({containerHeight:e,containerWidth:t,offset:i,targetHeight:n,targetWidth:o})=>{let r=u({height:n,offset:i,width:o}),s=o>n?o/t:n/e;return r>1?s:r*s},g=({containerHeight:e,containerWidth:t,hasScalableSrc:i,offset:n,targetHeight:o,targetWidth:r})=>e&&t?!i&&o&&r?m({containerHeight:e,containerWidth:t,offset:n,targetHeight:o,targetWidth:r}):u({height:e,offset:n,width:t}):1,f=/url(?:\(['"]?)(.*?)(?:['"]?\))/,p=e=>{if(e){if(h(e))return e.currentSrc;if(l(e)){let t=window.getComputedStyle(e).backgroundImage;if(t)return f.exec(t)?.[1]}}},v=e=>{if(e)return h(e)?e.alt??void 0:e.getAttribute("aria-label")??void 0},w=({containerHeight:e,containerLeft:t,containerTop:i,containerWidth:n,hasScalableSrc:o,offset:r,targetHeight:s,targetWidth:a})=>{let l=g({containerHeight:e,containerWidth:n,hasScalableSrc:o,offset:r,targetHeight:s,targetWidth:a});return{top:i,left:t,width:n*l,height:e*l,transform:`translate(0,0) scale(${1/l})`}},E=({position:e,relativeNum:t})=>{let i=parseFloat(e);return e.endsWith("%")?t*i/100:i},b=({containerHeight:e,containerLeft:t,containerTop:i,containerWidth:n,hasScalableSrc:o,objectFit:r,objectPosition:s,offset:a,targetHeight:l,targetWidth:h})=>{if("scale-down"===r&&(r=h<=n&&l<=e?"none":"contain"),"cover"===r||"contain"===r){let d=n/h,c=e/l,u="cover"===r?Math.max(d,c):Math.min(d,c),[m="50%",f="50%"]=s.split(" "),p=E({position:m,relativeNum:n-h*u}),v=E({position:f,relativeNum:e-l*u}),w=g({containerHeight:l*u,containerWidth:h*u,hasScalableSrc:o,offset:a,targetHeight:l,targetWidth:h});return{top:i+v,left:t+p,width:h*u*w,height:l*u*w,transform:`translate(0,0) scale(${1/w})`}}if("none"===r){let[r="50%",d="50%"]=s.split(" "),c=E({position:r,relativeNum:n-h}),u=E({position:d,relativeNum:e-l}),m=g({containerHeight:l,containerWidth:h,hasScalableSrc:o,offset:a,targetHeight:l,targetWidth:h});return{top:i+u,left:t+c,width:h*m,height:l*m,transform:`translate(0,0) scale(${1/m})`}}if("fill"!==r)return{};{let t=Math.max(n/h,e/l),i=g({containerHeight:l*t,containerWidth:h*t,hasScalableSrc:o,offset:a,targetHeight:l,targetWidth:h});return{width:n*i,height:e*i,transform:`translate(0,0) scale(${1/i})`}}},y=({backgroundPosition:e,backgroundSize:t,containerHeight:i,containerLeft:n,containerTop:o,containerWidth:r,hasScalableSrc:s,offset:a,targetHeight:l,targetWidth:h})=>{if("cover"===t||"contain"===t){let d=r/h,c=i/l,u="cover"===t?Math.max(d,c):Math.min(d,c),[m="50%",f="50%"]=e.split(" "),p=E({position:m,relativeNum:r-h*u}),v=E({position:f,relativeNum:i-l*u}),w=g({containerHeight:l*u,containerWidth:h*u,hasScalableSrc:s,offset:a,targetHeight:l,targetWidth:h});return{top:o+v,left:n+p,width:h*u*w,height:l*u*w,transform:`translate(0,0) scale(${1/w})`}}if("auto"===t){let[t="50%",d="50%"]=e.split(" "),c=E({position:t,relativeNum:r-h}),u=E({position:d,relativeNum:i-l}),m=g({containerHeight:l,containerWidth:h,hasScalableSrc:s,offset:a,targetHeight:l,targetWidth:h});return{top:o+u,left:n+c,width:h*m,height:l*m,transform:`translate(0,0) scale(${1/m})`}}{let[d="50%",c="50%"]=t.split(" "),u=Math.min(E({position:d,relativeNum:r})/h,E({position:c,relativeNum:i})/l),[m="50%",f="50%"]=e.split(" "),p=E({position:m,relativeNum:r-h*u}),v=E({position:f,relativeNum:i-l*u}),w=g({containerHeight:l*u,containerWidth:h*u,hasScalableSrc:s,offset:a,targetHeight:l,targetWidth:h});return{top:o+v,left:n+p,width:h*u*w,height:l*u*w,transform:`translate(0,0) scale(${1/w})`}}},S=/\.svg$/i,L=({hasZoomImg:e,imgSrc:t,isSvg:i,isZoomed:n,loadedImgEl:o,offset:r,shouldRefresh:s,targetEl:a})=>{let h=i||t?.slice?.(0,18)==="data:image/svg+xml"||e||!!(t&&S.test(t)),d=a.getBoundingClientRect(),c=window.getComputedStyle(a),u=null!=o&&l(a),m=Object.assign({},w({containerHeight:d.height,containerLeft:d.left,containerTop:d.top,containerWidth:d.width,hasScalableSrc:h,offset:r,targetHeight:o?.naturalHeight||d.height,targetWidth:o?.naturalWidth||d.width}),null==o||u?void 0:b({containerHeight:d.height,containerLeft:d.left,containerTop:d.top,containerWidth:d.width,hasScalableSrc:h,objectFit:c.objectFit,objectPosition:c.objectPosition,offset:r,targetHeight:o?.naturalHeight||d.height,targetWidth:o?.naturalWidth||d.width}),u?y({backgroundPosition:c.backgroundPosition,backgroundSize:c.backgroundSize,containerHeight:d.height,containerLeft:d.left,containerTop:d.top,containerWidth:d.width,hasScalableSrc:h,offset:r,targetHeight:o?.naturalHeight||d.height,targetWidth:o?.naturalWidth||d.width}):void 0);if(n){let e=window.innerWidth/2,t=window.innerHeight/2,i=parseFloat(String(m.left||0))+parseFloat(String(m.width||0))/2,n=parseFloat(String(m.top||0))+parseFloat(String(m.height||0))/2;s&&(m.transitionDuration="0.01ms"),m.transform=`translate(${e-i}px,${t-n}px) scale(1)`}return m},z=e=>{if(!e)return{};if(!c(e))return{height:e.offsetHeight,left:e.offsetLeft,width:e.offsetWidth,top:e.offsetTop};{let t=e.parentElement,i=e.getBoundingClientRect();if(!t)return{height:i.height,left:i.left,width:i.width,top:i.top};{let e=t.getBoundingClientRect();return{height:i.height,left:e.left-i.left,top:e.top-i.top,width:i.width}}}},C=e=>{let t="-zoom",i=["clip-path","fill","mask","marker-start","marker-mid","marker-end"],n=new Map;if(e.hasAttribute("id")){let i=e.id,o=i+t;n.set(i,o),e.id=o}e.querySelectorAll("[id]").forEach(e=>{let i=e.id,o=i+t;n.set(i,o),e.id=o}),n.forEach((t,n)=>{let o=`url(#${n})`,r=`url(#${t})`,s=i.map(e=>`[${e}="${o}"]`).join(", ");e.querySelectorAll(s).forEach(e=>{i.forEach(t=>{e.getAttribute(t)===o&&e.setAttribute(t,r)})})}),e.querySelectorAll("style").forEach(e=>{n.forEach((t,i)=>{e.textContent&&(e.textContent=e.textContent.replaceAll(`#${i}`,`#${t}`))})})},I=["img","svg",'[role="img"]',"[data-zoom]"].map(e=>`${e}:not([aria-hidden="true"])`).join(","),M={overflow:"",width:""};function D(e){return r.createElement(x,{...e})}class x extends r.Component{constructor(){super(...arguments),this.state={id:"",isZoomImgLoaded:!1,loadedImgEl:void 0,modalState:"UNLOADED",shouldRefresh:!1,styleGhost:{}},this.refContent=r.createRef(),this.refDialog=r.createRef(),this.refModalContent=r.createRef(),this.refModalImg=r.createRef(),this.refWrap=r.createRef(),this.imgEl=null,this.isScaling=!1,this.prevBodyAttrs=M,this.styleModalImg={},this.handleModalStateChange=e=>{let{modalState:t}=this.state;"LOADING"!==e&&"LOADING"===t?(this.loadZoomImg(),window.addEventListener("resize",this.handleResize,{passive:!0}),window.addEventListener("touchstart",this.handleTouchStart,{passive:!0}),window.addEventListener("touchmove",this.handleTouchMove,{passive:!0}),window.addEventListener("touchend",this.handleTouchEnd,{passive:!0}),window.addEventListener("touchcancel",this.handleTouchCancel,{passive:!0}),document.addEventListener("keydown",this.handleKeyDown,!0)):"LOADED"!==e&&"LOADED"===t?window.addEventListener("wheel",this.handleWheel,{passive:!0}):"UNLOADING"!==e&&"UNLOADING"===t?(this.ensureImgTransitionEnd(),window.removeEventListener("wheel",this.handleWheel),window.removeEventListener("touchstart",this.handleTouchStart),window.removeEventListener("touchmove",this.handleTouchMove),window.removeEventListener("touchend",this.handleTouchEnd),window.removeEventListener("touchcancel",this.handleTouchCancel),document.removeEventListener("keydown",this.handleKeyDown,!0)):"UNLOADED"!==e&&"UNLOADED"===t&&(this.bodyScrollEnable(),window.removeEventListener("resize",this.handleResize),this.refModalImg.current?.removeEventListener?.("transitionend",this.handleImgTransitionEnd),this.refDialog.current?.close?.())},this.getDialogContainer=()=>{let e=document.querySelector("[data-rmiz-portal]");return null==e&&((e=document.createElement("div")).setAttribute("data-rmiz-portal",""),document.body.appendChild(e)),e},this.setId=()=>{let e=()=>Math.random().toString(16).slice(-4);this.setState({id:e()+e()+e()})},this.setAndTrackImg=()=>{let e=this.refContent.current;!e||(this.imgEl=e.querySelector(I),this.imgEl?(this.contentNotFoundChangeObserver?.disconnect?.(),this.imgEl.addEventListener("load",this.handleImgLoad),this.imgEl.addEventListener("click",this.handleZoom),this.state.loadedImgEl||this.handleImgLoad(),this.imgElResizeObserver=new ResizeObserver(e=>{let t=e[0];t?.target&&(this.imgEl=t.target,this.setState({styleGhost:z(this.imgEl)}))}),this.imgElResizeObserver.observe(this.imgEl),this.contentChangeObserver||(this.contentChangeObserver=new MutationObserver(()=>{this.setState({styleGhost:z(this.imgEl)})}),this.contentChangeObserver.observe(e,{attributes:!0,childList:!0,subtree:!0}))):this.contentNotFoundChangeObserver||(this.contentNotFoundChangeObserver=new MutationObserver(this.setAndTrackImg),this.contentNotFoundChangeObserver.observe(e,{childList:!0,subtree:!0})))},this.handleIfZoomChanged=e=>{let{isZoomed:t}=this.props;!e&&t?this.zoom():e&&!t&&this.unzoom()},this.handleImgLoad=()=>{let e=p(this.imgEl);if(!e)return;let t=new Image;h(this.imgEl)&&(t.sizes=this.imgEl.sizes,t.srcset=this.imgEl.srcset,t.crossOrigin=this.imgEl.crossOrigin),t.src=e;let i=()=>{this.setState({loadedImgEl:t,styleGhost:z(this.imgEl)})};t.decode().then(i).catch(()=>{if(d(t)){i();return}t.onload=i})},this.handleZoom=()=>{this.hasImage()&&this.props.onZoomChange?.(!0)},this.handleUnzoom=()=>{this.props.onZoomChange?.(!1)},this.handleBtnUnzoomClick=e=>{e.preventDefault(),e.stopPropagation(),this.handleUnzoom()},this.handleDialogCancel=e=>{e.preventDefault()},this.handleDialogClick=e=>{(e.target===this.refModalContent.current||e.target===this.refModalImg.current)&&(e.stopPropagation(),this.handleUnzoom())},this.handleDialogClose=e=>{e.stopPropagation(),this.handleUnzoom()},this.handleKeyDown=e=>{("Escape"===e.key||27===e.keyCode)&&(e.preventDefault(),e.stopPropagation(),this.handleUnzoom())},this.handleWheel=e=>{e.ctrlKey||(e.stopPropagation(),queueMicrotask(()=>{this.handleUnzoom()}))},this.handleTouchStart=e=>{if(e.touches.length>1){this.isScaling=!0;return}1===e.changedTouches.length&&e.changedTouches[0]&&(this.touchYStart=e.changedTouches[0].screenY)},this.handleTouchMove=e=>{let t=window.visualViewport?.scale??1;this.props.canSwipeToUnzoom&&!this.isScaling&&t<=1&&null!=this.touchYStart&&e.changedTouches[0]&&(this.touchYEnd=e.changedTouches[0].screenY,Math.abs(Math.max(this.touchYStart,this.touchYEnd)-Math.min(this.touchYStart,this.touchYEnd))>this.props.swipeToUnzoomThreshold&&(this.touchYStart=void 0,this.touchYEnd=void 0,this.handleUnzoom()))},this.handleTouchEnd=()=>{this.isScaling=!1,this.touchYStart=void 0,this.touchYEnd=void 0},this.handleTouchCancel=()=>{this.isScaling=!1,this.touchYStart=void 0,this.touchYEnd=void 0},this.handleResize=()=>{this.setState({shouldRefresh:!0})},this.hasImage=()=>this.imgEl&&(this.state.loadedImgEl||c(this.imgEl))&&"none"!==window.getComputedStyle(this.imgEl).display,this.zoom=()=>{this.bodyScrollDisable(),this.refDialog.current?.showModal?.(),this.refModalImg.current?.addEventListener?.("transitionend",this.handleImgTransitionEnd),this.setState({modalState:"LOADING"})},this.unzoom=()=>{this.setState({modalState:"UNLOADING"})},this.handleImgTransitionEnd=()=>{clearTimeout(this.timeoutTransitionEnd),"LOADING"===this.state.modalState?this.setState({modalState:"LOADED"}):"UNLOADING"===this.state.modalState&&this.setState({shouldRefresh:!1,modalState:"UNLOADED"})},this.ensureImgTransitionEnd=()=>{if(this.refModalImg.current){let e=window.getComputedStyle(this.refModalImg.current).transitionDuration,t=parseFloat(e);if(t){let i=t*(e.endsWith("ms")?1:1e3)+50;this.timeoutTransitionEnd=setTimeout(this.handleImgTransitionEnd,i)}}},this.bodyScrollDisable=()=>{this.prevBodyAttrs={overflow:document.body.style.overflow,width:document.body.style.width};let e=document.body.clientWidth;document.body.style.overflow="hidden",document.body.style.width=`${e}px`},this.bodyScrollEnable=()=>{document.body.style.width=this.prevBodyAttrs.width,document.body.style.overflow=this.prevBodyAttrs.overflow,this.prevBodyAttrs=M},this.loadZoomImg=()=>{let{props:{zoomImg:e}}=this,t=e?.src;if(t){let i=new Image;i.sizes=e?.sizes??"",i.srcset=e?.srcSet??"",i.crossOrigin=e?.crossOrigin??void 0,i.src=t;let n=()=>{this.setState({isZoomImgLoaded:!0})};i.decode().then(n).catch(()=>{if(d(i)){n();return}i.onload=n})}},this.UNSAFE_handleSvg=()=>{let{imgEl:e,refModalImg:t,styleModalImg:i}=this;if(c(e)){let n=e.cloneNode(!0);C(n),n.style.width=`${i.width||0}px`,n.style.height=`${i.height||0}px`,n.addEventListener("click",this.handleUnzoom),t.current?.firstChild?.remove?.(),t.current?.appendChild?.(n)}}}render(){let{handleBtnUnzoomClick:e,handleDialogCancel:t,handleDialogClick:i,handleDialogClose:n,handleUnzoom:o,handleZoom:a,imgEl:d,props:{a11yNameButtonUnzoom:u,a11yNameButtonZoom:m,children:g,classDialog:f,IconUnzoom:w,IconZoom:E,isZoomed:b,wrapElement:y,ZoomContent:S,zoomImg:z,zoomMargin:C},refContent:I,refDialog:M,refModalContent:D,refModalImg:x,refWrap:O,state:{id:N,isZoomImgLoaded:A,loadedImgEl:j,modalState:T,shouldRefresh:P,styleGhost:k}}=this,R=`rmiz-modal-${N}`,U=`rmiz-modal-img-${N}`,_=l(d),W=h(d),$=c(d),F=v(d),G=p(d),B=W?d.sizes:void 0,H=W?d.srcset:void 0,Z=W?d.crossOrigin:void 0,Y=!!z?.src,q=this.hasImage(),V=F?`${m}: ${F}`:m,K="LOADING"===T||"LOADED"===T,J=q?"found":"not-found",Q="UNLOADED"===T||"UNLOADING"===T?"hidden":"visible";this.styleModalImg=q?L({hasZoomImg:Y,imgSrc:G,isSvg:$,isZoomed:b&&K,loadedImgEl:j,offset:C,shouldRefresh:P,targetEl:d}):{};let X=null;if(q){let t=W||_?r.createElement("img",{alt:F,crossOrigin:Z,sizes:B,src:G,srcSet:H,...A&&"LOADED"===T?z:{},"data-rmiz-modal-img":"",height:this.styleModalImg.height||void 0,id:U,ref:x,style:this.styleModalImg,width:this.styleModalImg.width||void 0}):$?r.createElement("div",{"data-rmiz-modal-img":!0,ref:x,style:this.styleModalImg}):null,i=r.createElement("button",{"aria-label":u,"data-rmiz-btn-unzoom":"",onClick:e,type:"button"},r.createElement(w,null));X=S?r.createElement(S,{buttonUnzoom:i,modalState:T,img:t,onUnzoom:o}):r.createElement(r.Fragment,null,t,i)}return r.createElement(y,{"aria-owns":R,"data-rmiz":"",ref:O},r.createElement(y,{"data-rmiz-content":J,ref:I,style:{visibility:"UNLOADED"===T?"visible":"hidden"}},g),q&&r.createElement(y,{"data-rmiz-ghost":"",style:k},r.createElement("button",{"aria-label":V,"data-rmiz-btn-zoom":"",onClick:a,type:"button"},r.createElement(E,null))),q&&s.createPortal(r.createElement("dialog",{"aria-labelledby":U,"aria-modal":"true",className:f,"data-rmiz-modal":"",id:R,onClick:i,onClose:n,onCancel:t,ref:M,role:"dialog"},r.createElement("div",{"data-rmiz-modal-overlay":Q}),r.createElement("div",{"data-rmiz-modal-content":"",ref:D},X)),this.getDialogContainer()))}componentDidMount(){this.setId(),this.setAndTrackImg(),this.handleImgLoad(),this.UNSAFE_handleSvg()}componentWillUnmount(){"UNLOADED"!==this.state.modalState&&this.bodyScrollEnable(),this.contentChangeObserver?.disconnect?.(),this.contentNotFoundChangeObserver?.disconnect?.(),this.imgElResizeObserver?.disconnect?.(),this.imgEl?.removeEventListener?.("load",this.handleImgLoad),this.imgEl?.removeEventListener?.("click",this.handleZoom),this.refModalImg.current?.removeEventListener?.("transitionend",this.handleImgTransitionEnd),window.removeEventListener("wheel",this.handleWheel),window.removeEventListener("touchstart",this.handleTouchStart),window.removeEventListener("touchmove",this.handleTouchMove),window.removeEventListener("touchend",this.handleTouchEnd),window.removeEventListener("touchcancel",this.handleTouchCancel),window.removeEventListener("resize",this.handleResize),document.removeEventListener("keydown",this.handleKeyDown,!0)}componentDidUpdate(e,t){this.handleModalStateChange(t.modalState),this.UNSAFE_handleSvg(),this.handleIfZoomChanged(e.isZoomed)}}function O(e){let[t,i]=r.useState(!1);return r.createElement(D,{...e,isZoomed:t,onZoomChange:i})}x.defaultProps={a11yNameButtonUnzoom:"Minimize image",a11yNameButtonZoom:"Expand image",canSwipeToUnzoom:!0,IconUnzoom:function(){return r.createElement("svg",{"aria-hidden":"true","data-rmiz-btn-unzoom-icon":!0,fill:"currentColor",focusable:"false",viewBox:"0 0 16 16",xmlns:"http://www.w3.org/2000/svg"},r.createElement("path",{d:"M 14.144531 1.148438 L 9 6.292969 L 9 3 L 8 3 L 8 8 L 13 8 L 13 7 L 9.707031 7 L 14.855469 1.851563 Z M 8 8 L 3 8 L 3 9 L 6.292969 9 L 1.148438 14.144531 L 1.851563 14.855469 L 7 9.707031 L 7 13 L 8 13 Z"}))},IconZoom:function(){return r.createElement("svg",{"aria-hidden":"true","data-rmiz-btn-zoom-icon":!0,fill:"currentColor",focusable:"false",viewBox:"0 0 16 16",xmlns:"http://www.w3.org/2000/svg"},r.createElement("path",{d:"M 9 1 L 9 2 L 12.292969 2 L 2 12.292969 L 2 9 L 1 9 L 1 14 L 6 14 L 6 13 L 2.707031 13 L 13 2.707031 L 13 6 L 14 6 L 14 1 Z"}))},swipeToUnzoomThreshold:10,wrapElement:"div",zoomMargin:0};var N=i(5675),A=i.n(N);let j=(0,r.forwardRef)((e,t)=>{let i="object"==typeof e.src?A():"img";return(0,o.jsx)(i,{...e,ref:t})});j.displayName="Image";let T={img:e=>{var t;let i=(0,r.useRef)(null),[n,s]=(0,r.useState)(!1);(0,r.useEffect)(()=>{s(null!==i.current.closest("a"))},[]);let a=(0,o.jsx)(j,{...e,ref:i});return n?a:(0,o.jsx)(O,{zoomMargin:40,zoomImg:{src:"string"==typeof(t=e.src)?t:"default"in t?t.default.src:t.src,alt:e.alt},wrapElement:"span",children:a})}},P=e=>({...T,...(0,n.a)(e)})},7812:function(e,t,i){"use strict";i.d(t,{c:function(){return h}});var n=i(5893),o=i(3665),r=i(8644);let s=(0,i(7294).createContext)({}),a=s.Provider;s.displayName="SSG";var l=i(9299);function h(e,t,i,n){let r=globalThis[o.ud];return r.route=t,r.pageMap=i.pageMap,r.context[t]={Content:e,pageOpts:i,useTOC:n},d}function d({__nextra_pageMap:e=[],__nextra_dynamic_opts:t,...i}){let s=globalThis[o.ud],{Layout:l,themeConfig:h}=s,{route:d,locale:u}=(0,r.t)(),m=s.context[d];if(!m)throw Error(`No content found for the "${d}" route. Please report it as a bug.`);let{pageOpts:g,useTOC:f,Content:p}=m;if(d.startsWith("/["))g.pageMap=e;else for(let{route:t,children:i}of e){let e=t.split("/").slice(u?2:1);(function e(t,[i,...n]){for(let o of t)if("children"in o&&i===o.name)return n.length?e(o.children,n):o})(g.pageMap,e).children=i}if(t){let{title:e,frontMatter:i}=t;g={...g,title:e,frontMatter:i}}return(0,n.jsx)(l,{themeConfig:h,pageOpts:g,pageProps:i,children:(0,n.jsx)(a,{value:i,children:(0,n.jsx)(c,{useTOC:f,children:(0,n.jsx)(p,{...i})})})})}function c({children:e,useTOC:t}){let{wrapper:i}=(0,l.a)();return(0,n.jsx)(u,{useTOC:t,wrapper:i,children:e})}function u({children:e,useTOC:t,wrapper:i,...o}){let r=t(o);return i?(0,n.jsx)(i,{toc:r,children:e}):e}}}]);