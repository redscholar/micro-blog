/*
  This example requires Tailwind CSS v2.0+ 
  
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/forms'),
    ],
  }
  ```
*/
import { LockClosedIcon } from '@heroicons/react/20/solid';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { postSignIn } from '../../api/sign';

export const SignIn = (props: any) => {
  const [err,setErr] = useState("")
  const [data, setData] = useState({
    username: "",
    password: ""
  })
  const signIn = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // 调用接口
    postSignIn(data).then((res) => {
      if (res.code == 0) {
        localStorage.setItem("AccessToken", res.data)
        window.location.href = "/"
      } else {
        setErr(res.msg)
      }
    }).catch(() => {
      setErr("登录失败")
    })
  }
  

  return (
    <div className="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="w-full max-w-md space-y-8">
        <div>
          <img
            className="mx-auto h-12 w-auto"
            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
            alt="Your Company"
          />
        </div>
        <form className="mt-8 space-y-6" method="POST" onSubmit={signIn}>
          <input type="hidden" name="remember" defaultValue="true" />
          <div className="-space-y-px rounded-md shadow-sm">
            <div>
              <label htmlFor="username" className="sr-only">
                用户名
              </label>
              <input
                id="username"
                name="username"
                type="text"
                onChange={(t) => {
                  setErr("")
                  setData((d) => ({ ...d, username: t.target.value }))
                }}
                required
                className="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                placeholder="用户名"
              />
            </div>
            <div>
              <label htmlFor="password" className="sr-only">
                密码
              </label>
              <input
                id="password"
                name="password"
                type="password"
                onChange={(t) => {
                  setErr("")
                  setData((d) => ({ ...d, password: t.target.value }))
                }}
                required
                className="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                placeholder="密码"
              />
            </div>
          </div>

          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <input
                id="remember-me"
                name="remember-me"
                type="checkbox"
                className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
              />
              <label htmlFor="remember-me" className="ml-2 block text-sm text-gray-900">
                记住我
              </label>
            </div>

            <div className="text-sm">
              <Link to="/signUp" className="font-medium text-indigo-600 hover:text-indigo-500">
                没有账号？前去注册！
              </Link>
            </div>
          </div>

          <div>
            <button
              type="submit"
              className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
              <span className="absolute inset-y-0 left-0 flex items-center pl-3">
                <LockClosedIcon className="h-5 w-5 text-indigo-500 group-hover:text-indigo-400" aria-hidden="true" />
              </span>
              登录
            </button>
            <p className="relative flex w-full justify-center rounded-md border border-transparent text-red-600 break-words">
              {err}
            </p>
          </div>
        </form>
      </div>
    </div>
  )
}