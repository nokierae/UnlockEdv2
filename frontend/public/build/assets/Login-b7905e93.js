import{j as e,W as u,a as x}from"./app-6ded1f35.js";import{G as p}from"./GuestLayout-bcedc2b3.js";import{I as n,T as o,a as l}from"./TextInput-699a9a09.js";import{P as g}from"./PrimaryButton-14133a1e.js";function j({className:r="",...a}){return e.jsx("input",{...a,type:"checkbox",className:"rounded dark:bg-gray-900 border-gray-300 dark:border-gray-700 text-indigo-600 shadow-sm focus:ring-indigo-500 dark:focus:ring-indigo-600 dark:focus:ring-offset-gray-800 "+r})}function N({status:r}){const{data:a,setData:t,post:d,processing:i,errors:m}=u({username:"",password:"",remember:!1}),c=s=>{s.preventDefault(),d(route("login"))};return e.jsxs(p,{children:[e.jsx(x,{title:"Log in"}),r&&e.jsx("div",{className:"mb-4 font-medium text-sm text-green-600",children:r}),e.jsxs("form",{onSubmit:c,children:[e.jsxs("div",{className:"mt-4",children:[e.jsx(n,{htmlFor:"username",value:"Username"}),e.jsx(o,{id:"username",type:"username",name:"username",value:a.username,className:"mt-1 block w-full",autoComplete:"username",onChange:s=>t("username",s.target.value)}),e.jsx(l,{message:m.username,className:"mt-2"})]}),e.jsxs("div",{className:"mt-4",children:[e.jsx(n,{htmlFor:"password",value:"Password"}),e.jsx(o,{id:"password",type:"password",name:"password",value:a.password,className:"mt-1 block w-full",autoComplete:"current-password",onChange:s=>t("password",s.target.value)}),e.jsx(l,{message:m.password,className:"mt-2"})]}),e.jsx("div",{className:"block mt-4",children:e.jsxs("label",{className:"flex items-center",children:[e.jsx(j,{name:"remember",checked:a.remember,onChange:s=>t("remember",s.target.checked)}),e.jsx("span",{className:"ms-2 text-sm text-slate-600 dark:text-slate-400",children:"Remember me"})]})}),e.jsx("div",{className:"flex items-center justify-end mt-4",children:e.jsx(g,{className:"ms-4",disabled:i,children:"Log in"})})]})]})}export{N as default};