/* This example requires Tailwind CSS v2.0+ */
import { Dialog, Menu, Transition } from '@headlessui/react'
import { BellIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { Dispatch, FormEventHandler, Fragment, SetStateAction, useEffect, useState } from 'react'
import { Link } from "react-router-dom"
import { classNames } from '.'
import { postChangePwd } from '../../../api/sign'
import { MessageData, showMessage } from '../../component/message'
import { Modal } from '../../component/modal'
export const ToolBar = () => {
    const [changePwd, setChangePwd] = useState(false)

    // 查询用户信息
    return (
        <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
            <button
                type="button"
                className="rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
            >
                <span className="sr-only">消息通知</span>
                <BellIcon className="h-6 w-6" aria-hidden="true" />
            </button>

            {/* Profile dropdown */}
            <Menu as="div" className="relative ml-3">
                <div>
                    <Menu.Button className="flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
                        <span className="sr-only">我的菜单</span>
                        <img
                            className="h-8 w-8 rounded-full"
                            src="image/user.png"
                            alt=""
                        />
                    </Menu.Button>
                </div>
                <Transition
                    as={Fragment}
                    enter="transition ease-out duration-100"
                    enterFrom="transform opacity-0 scale-95"
                    enterTo="transform opacity-100 scale-100"
                    leave="transition ease-in duration-75"
                    leaveFrom="transform opacity-100 scale-100"
                    leaveTo="transform opacity-0 scale-95"
                >
                    <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                        <Menu.Item>
                            {({ active }) => (
                                <a
                                    onClick={() => { setChangePwd(true) }}
                                    className={classNames(active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700')}
                                >
                                    修改密码
                                </a>
                            )}
                        </Menu.Item>
                        <Menu.Item>
                            {({ active }) => (
                                <Link
                                    to="/signOut"
                                    className={classNames(active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700')}
                                >
                                    登出
                                </Link>
                            )}
                        </Menu.Item>
                    </Menu.Items>
                </Transition>
            </Menu>
            <Modal open={changePwd} setOpen={setChangePwd} child={ChangePwdForm()} />
            <Bulletin />
        </div>

    )
}


// 修改密码
const ChangePwdForm = () => {
    const [formdata, setformdata] = useState({
        oldPwd: "",
        newPwd: ""
    })
    const changePwd = (e: React.FormEvent<HTMLFormElement>) => {
        console.log("changepwd", e)
        e.preventDefault();
        postChangePwd({ oldPwd: formdata.oldPwd, newPwd: formdata.newPwd }).then(() => {
            showMessage(new MessageData(true, "success", "", "修改成功"))
        })
    }
    
    return (
        <div className="mt-5 md:col-span-3 md:mt-3">
            <form method="POST" onSubmit={changePwd}>
                <div className="overflow-hidden shadow sm:rounded-md mt-10">
                    <div className="bg-white px-4 py-5 sm:p-6">
                        <div className="grid grid-cols-3 gap-6">
                            <div className="col-span-1 sm:col-span-1 block text-sm font-medium text-gray-700 items-center center">
                                {/* <label htmlFor="oldPassword" className="block text-sm font-medium text-gray-700"> */}
                                原密码
                            </div>
                            <div className="col-span-2 sm:col-span-2">

                                <input
                                    type="password"
                                    name="oldPassword"
                                    id="oldPassword"
                                    onChange={(t) => setformdata((f) => ({ ...f, oldPwd: t.target.value }))}
                                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                />
                            </div>

                            <div className="col-span-1 sm:col-span-1 block text-sm font-medium text-gray-700 items-center">
                                {/* <label htmlFor="newPassword" className="block text-sm font-medium text-gray-700"> */}
                                新密码
                                {/* </label> */}
                            </div>
                            <div className="col-span-2 sm:col-span-2">

                                <input
                                    type="password"
                                    name="newPassword"
                                    id="newPassword"
                                    onChange={(t) => setformdata((f) => ({ ...f, newPwd: t.target.value }))}
                                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                />
                            </div>
                        </div>
                    </div>
                    
                    <div className="bg-gray-50 px-4 py-3 text-right sm:px-6">
                        <button
                            type="submit"
                            className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                        >
                            提交
                        </button>
                    </div>
                </div>
            </form>
        </div>
    )
}

// 默认公告
const bulletin = {
    name: '感谢授权 ',
    text: '感谢王总授权，原始项目地址：https://github.com/xiaowang012/gin-blog',
    imageSrc: 'image/auth.png',
    imageAlt: '',

}

const Bulletin = () => {
    const [open, setOpen] = useState(false)
    return (
        <>
            <button type="button" onClick={() => setOpen(true)}
                className="ml-3 float-right align-baseline rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
            >
                <span className="sr-only">查看公告</span>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M12 18v-5.25m0 0a6.01 6.01 0 001.5-.189m-1.5.189a6.01 6.01 0 01-1.5-.189m3.75 7.478a12.06 12.06 0 01-4.5 0m3.75 2.383a14.406 14.406 0 01-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 10-7.517 0c.85.493 1.509 1.333 1.509 2.316V18" />
                </svg>
            </button>

            <Modal open={open} setOpen={setOpen} child={
                <div className="grid w-full grid-cols-1 items-start gap-y-8 gap-x-6 sm:grid-cols-12 lg:gap-x-8">
                    <div className="aspect-w-2 aspect-h-3 overflow-hidden rounded-lg bg-gray-100 sm:col-span-4 lg:col-span-5">
                        <img src={bulletin.imageSrc} alt={bulletin.imageAlt} className="object-cover object-center" />
                    </div>
                    <div className="sm:col-span-8 lg:col-span-7">
                        <h2 className="text-2xl font-bold text-gray-900 sm:pr-12">{bulletin.name}</h2>
                        <p className="mt-4">{bulletin.text}</p>
                    </div>
                </div>
            } />
        </>

    )
}
