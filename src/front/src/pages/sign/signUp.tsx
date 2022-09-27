import { useState } from "react";
import { Link } from "react-router-dom";
import { postSignUp } from "../../api/sign";

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
export default function SignUp() {
    const [data, setData] = useState({
        username: "",
        password: ""
    })
    const [err, setErr] = useState("")

    const signUp = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        postSignUp(data).then((res) => {
            if (res.code == 0) {
                localStorage.setItem("AccessToken", res.data)
                window.location.href = "/"
            } else {
                setErr(res.msg)
            }
        }).catch(() => {
            setErr("注册失败")
        })

    }

    return (
        <div className="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="w-full max-w-md space-y-8">
                <div>
                    <h2 className="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">
                        注册账号
                    </h2>
                </div>
                <form className="mt-8 space-y-6" method="POST" onSubmit={signUp}>
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
                        <div className="text-sm">
                            <Link to="/signIn" className="font-medium text-indigo-600 hover:text-indigo-500">
                                已有账号，前去登录！
                            </Link>
                        </div>
                    </div>

                    <div>
                        <button
                            type="submit"
                            className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                        >
                            确定
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
