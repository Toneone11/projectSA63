(window.webpackJsonp=window.webpackJsonp||[]).push([[4],{463:function(e,t,a){a(464),e.exports=a(913)},913:function(e,t,a){"use strict";a.r(t);var n={};a.r(n),a.d(n,"WelcomePlugin",(function(){return W}));a(467);var r=a(0),i=a.n(r),s=a(17),o=a.n(s),c=a(43),l=a(64),u=a(392),d=a(400),m=a(388),h=a(382),p=a(967),f=a(397),w=a(389),g=a(99),y=a(393);const E="http://localhost:8080/api/v1".replace(/\/+$/,"");class b{constructor(e=new U){this.configuration=e,b.prototype.__init.call(this),this.middleware=e.middleware}withMiddleware(...e){const t=this.clone();return t.middleware=t.middleware.concat(...e),t}withPreMiddleware(...e){const t=e.map(e=>({pre:e}));return this.withMiddleware(...t)}withPostMiddleware(...e){const t=e.map(e=>({post:e}));return this.withMiddleware(...t)}async request(e){const{url:t,init:a}=this.createFetchParams(e),n=await this.fetchApi(t,a);if(n.status>=200&&n.status<300)return n;throw n}createFetchParams(e){let t=this.configuration.basePath+e.path;void 0!==e.query&&0!==Object.keys(e.query).length&&(t+="?"+this.configuration.queryParamsStringify(e.query));const a="undefined"!=typeof FormData&&e.body instanceof FormData||e.body instanceof URLSearchParams||(n=e.body,"undefined"!=typeof Blob&&n instanceof Blob)?e.body:JSON.stringify(e.body);var n;const r=Object.assign({},this.configuration.headers,e.headers);return{url:t,init:{method:e.method,headers:r,body:a,credentials:this.configuration.credentials}}}__init(){this.fetchApi=async(e,t)=>{let a={url:e,init:t};for(const e of this.middleware)e.pre&&(a=await e.pre({fetch:this.fetchApi,...a})||a);let n=await this.configuration.fetchApi(a.url,a.init);for(const a of this.middleware)a.post&&(n=await a.post({fetch:this.fetchApi,url:e,init:t,response:n.clone()})||n);return n}}clone(){const e=new(0,this.constructor)(this.configuration);return e.middleware=this.middleware.slice(),e}}class v extends Error{__init2(){this.name="RequiredError"}constructor(e,t){super(t),this.field=e,v.prototype.__init2.call(this)}}class U{constructor(e={}){this.configuration=e}get basePath(){return null!=this.configuration.basePath?this.configuration.basePath:E}get fetchApi(){return this.configuration.fetchApi||window.fetch.bind(window)}get middleware(){return this.configuration.middleware||[]}get queryParamsStringify(){return this.configuration.queryParamsStringify||q}get username(){return this.configuration.username}get password(){return this.configuration.password}get apiKey(){const e=this.configuration.apiKey;if(e)return"function"==typeof e?e:()=>e}get accessToken(){const e=this.configuration.accessToken;if(e)return"function"==typeof e?e:()=>e}get headers(){return this.configuration.headers}get credentials(){return this.configuration.credentials}}function R(e,t){const a=e[t];return null!=a}function q(e,t=""){return Object.keys(e).map(a=>{const n=t+(t.length?`[${a}]`:a),r=e[a];if(r instanceof Array){const e=r.map(e=>encodeURIComponent(String(e))).join(`&${encodeURIComponent(n)}=`);return`${encodeURIComponent(n)}=${e}`}return r instanceof Object?q(r,n):`${encodeURIComponent(n)}=${encodeURIComponent(String(r))}`}).filter(e=>e.length>0).join("&")}class j{constructor(e,t=(e=>e)){this.raw=e,this.transformer=t}async value(){return this.transformer(await this.raw.json())}}function P(e){return function(e,t){if(null==e)return e;return{age:R(e,"age")?e.age:void 0,id:R(e,"id")?e.id:void 0,name:R(e,"name")?e.name:void 0}}(e)}function O(e){if(void 0!==e)return null===e?null:{age:e.age,id:e.id,name:e.name}}class S extends b{async createUserRaw(e){if(null===e.user||void 0===e.user)throw new v("user","Required parameter requestParameters.user was null or undefined when calling createUser.");const t={"Content-Type":"application/json"},a=await this.request({path:"/users",method:"POST",headers:t,query:{},body:O(e.user)});return new j(a,e=>P(e))}async createUser(e){const t=await this.createUserRaw(e);return await t.value()}async deleteUserRaw(e){if(null===e.id||void 0===e.id)throw new v("id","Required parameter requestParameters.id was null or undefined when calling deleteUser.");const t=await this.request({path:"/users/{id}".replace("{id}",encodeURIComponent(String(e.id))),method:"DELETE",headers:{},query:{}});return new j(t)}async deleteUser(e){const t=await this.deleteUserRaw(e);return await t.value()}async getUserRaw(e){if(null===e.id||void 0===e.id)throw new v("id","Required parameter requestParameters.id was null or undefined when calling getUser.");const t=await this.request({path:"/users/{id}".replace("{id}",encodeURIComponent(String(e.id))),method:"GET",headers:{},query:{}});return new j(t,e=>P(e))}async getUser(e){const t=await this.getUserRaw(e);return await t.value()}async listUserRaw(e){const t={};void 0!==e.limit&&(t.limit=e.limit),void 0!==e.offset&&(t.offset=e.offset);const a=await this.request({path:"/users",method:"GET",headers:{},query:t});return new j(a,e=>e.map(P))}async listUser(e){const t=await this.listUserRaw(e);return await t.value()}async updateUserRaw(e){if(null===e.id||void 0===e.id)throw new v("id","Required parameter requestParameters.id was null or undefined when calling updateUser.");if(null===e.user||void 0===e.user)throw new v("user","Required parameter requestParameters.user was null or undefined when calling updateUser.");const t={"Content-Type":"application/json"},a=await this.request({path:"/users/{id}".replace("{id}",encodeURIComponent(String(e.id))),method:"PUT",headers:t,query:{},body:O(e.user)});return new j(a,e=>P(e))}async updateUser(e){const t=await this.updateUserRaw(e);return await t.value()}}const T=Object(u.a)({table:{minWidth:650}});function _(){const e=T(),t=new S,[a,n]=Object(r.useState)(Array),[s,o]=Object(r.useState)(!0);Object(r.useEffect)(()=>{(async()=>{const e=await t.listUser({limit:10,offset:0});o(!1),n(e)})()},[s]);return i.a.createElement(p.a,{component:g.a},i.a.createElement(d.a,{className:e.table,"aria-label":"simple table"},i.a.createElement(f.a,null,i.a.createElement(w.a,null,i.a.createElement(h.a,{align:"center"},"ลำดับที่."),i.a.createElement(h.a,{align:"center"},"รหัสผู้เข้ารับการผ่าตัด"),i.a.createElement(h.a,{align:"center"},"ชื่อผู้เข้ารับการผ่าตัด"),i.a.createElement(h.a,{align:"center"},"ห้องผ่าตัด"),i.a.createElement(h.a,{align:"center"},"วันที่-เวลา"))),i.a.createElement(m.a,null,a.map(e=>i.a.createElement(w.a,{key:e.id},i.a.createElement(h.a,{align:"center"},e.id),i.a.createElement(h.a,{align:"center"},e.p_id),i.a.createElement(h.a,{align:"center"},e.patient),i.a.createElement(h.a,{align:"center"},e.room),i.a.createElement(h.a,{align:"center"},e.datetime),i.a.createElement(h.a,{align:"center"},i.a.createElement(y.a,{onClick:()=>{(async e=>{await t.deleteUser({id:e});o(!0)})(e.id)},style:{marginLeft:10},variant:"contained",color:"secondary"},"Delete")))))))}var C=()=>i.a.createElement(c.e,{theme:c.h.home},i.a.createElement(c.c,{title:"ยินดีต้อนรับเข้าสู่ ระบบจองห้องผ่าตัด",subtitle:"ผ่าแล้วเย็บคืนด้วย"}),i.a.createElement(c.a,null,i.a.createElement(c.b,{title:"ตารางการจองห้องผ่าตัด"},i.a.createElement("font",{color:"blue"},"นายแพทย์ ทดสอบ ทดสอบ"),i.a.createElement(y.a,{style:{marginLeft:20},component:l.b,color:"primary",to:"/",variant:"contained"},"ออกจากระบบ")),i.a.createElement(_,null),i.a.createElement("br",null),i.a.createElement(c.d,{component:l.b,to:"/user"},i.a.createElement(y.a,{variant:"contained",color:"primary"},"เพิ่มการจองห้อง")))),I=a(394),k=a(391),L=a(281),A=a(971),N=a(975);const x=Object(u.a)(e=>Object(I.a)({root:{display:"flex",flexWrap:"wrap",justifyContent:"center"},margin:{margin:e.spacing(3)},withoutLabel:{marginTop:e.spacing(3)},textField:{width:"25ch"}})),$={name:"",age:69};function F(){const e=x(),t=new S,[a,n]=Object(r.useState)($),[s,o]=Object(r.useState)(!1),[u,d]=Object(r.useState)(!0);return i.a.createElement(c.e,{theme:c.h.home},i.a.createElement(c.c,{title:"ในระบบจองห้องผ่าตัด แพทย์สามารถจองห้องผ่าตัดที่นี่",subtitle:"ผ่าแล้วเย็บคืนด้วย"}),i.a.createElement(c.a,null,i.a.createElement(c.b,{title:"กรอกรายละเอียดการจอง"},i.a.createElement("font",{color:"blue"},"นายแพทย์ ทดสอบ ทดสอบ"),i.a.createElement(y.a,{style:{marginLeft:20},component:l.b,color:"primary",to:"/",variant:"contained"},"ออกจากระบบ"),s?i.a.createElement("div",null,u?i.a.createElement(A.a,{severity:"success"},"This is a success alert — check it out!"):i.a.createElement(A.a,{severity:"warning",style:{marginTop:20}},"This is a warning alert — check it out!")):null),i.a.createElement("div",{className:e.root},i.a.createElement("form",{noValidate:!0,autoComplete:"off"},i.a.createElement(L.a,{fullWidth:!0,className:e.margin,variant:"outlined"},i.a.createElement(N.a,{id:"docter",options:[{doctor_id:"4001"},{doctor_id:"4002"}],getOptionLabel:e=>e.doctor_id,style:{width:300},renderInput:e=>i.a.createElement(k.a,{...e,label:"รหัสประจำตัวแพทย์",variant:"outlined"})})),i.a.createElement(L.a,{fullWidth:!0,className:e.margin,variant:"outlined"},i.a.createElement(N.a,{id:"patient",options:[{patient_id:"5001"},{patient_id:"5002"}],getOptionLabel:e=>e.patient_id,style:{width:300},renderInput:e=>i.a.createElement(k.a,{...e,label:"รหัสผู้เข้ารับการรักษา",variant:"outlined"})})),i.a.createElement(L.a,{fullWidth:!0,className:e.margin,variant:"outlined"},i.a.createElement(N.a,{id:"room",options:[{room:"6001"},{room:"6002"}],getOptionLabel:e=>e.room,style:{width:300},renderInput:e=>i.a.createElement(k.a,{...e,label:"ห้องผ่าตัด",variant:"outlined"})})),i.a.createElement("form",{className:e.margin,noValidate:!0},i.a.createElement(k.a,{id:"datetime-local",label:"วัน-เวลาที่ทำการผ่าตัด",type:"datetime-local",defaultValue:"2017-05-24T10:30",className:e.textField,InputLabelProps:{shrink:!0}})),i.a.createElement("div",{className:e.margin},i.a.createElement(y.a,{onClick:()=>{(async()=>{const e=await t.createUser({user:a});o(!0),""!=e.id?d(!0):d(!1);setTimeout(()=>{o(!1)},1e3)})()},variant:"contained",color:"primary"},"ยืนยันการจองห้อง"),i.a.createElement(y.a,{style:{marginLeft:20},component:l.b,to:"/",variant:"contained"},"ย้อนกลับ"))))))}const W=Object(c.g)({id:"welcome",register({router:e}){e.registerRoute("/",C),e.registerRoute("/user",F)}}),M=Object(c.f)({plugins:Object.values(n)}),D=M.getProvider(),B=M.getRouter(),J=M.getRoutes();var V=()=>i.a.createElement(D,null,i.a.createElement(B,null,i.a.createElement(J,null)));o.a.render(i.a.createElement(V,null),document.getElementById("root"))}},[[463,9,6,1,3,8,5,0,2,7,10]]]);
//# sourceMappingURL=main.90b320fe.chunk.js.map