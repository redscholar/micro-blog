import { Dialog, Transition } from '@headlessui/react'
import { CheckCircleIcon, ExclamationCircleIcon, XCircleIcon } from '@heroicons/react/20/solid'
import { Fragment, useEffect, useState } from "react"

import EventEmitter from 'events'


export class MessageData {
    use: boolean
    type: "error" | "warning" | "success" | undefined
    title: string
    content: string
    expire: number
    constructor(use?: boolean, type?: "error" | "warning" | "success", title?: string, content?: string, expire?: number) {
        this.use = use || false
        this.type = type || undefined
        this.title = title || ""
        this.content = content || ""
        this.expire = expire || 3000

    }
}
const messageEvent = new EventEmitter()

export const showMessage = (pub: MessageData) => {
    messageEvent.emit("message", pub)
}


export const MessageLayout = () => {
    const [data, setData] = useState(new MessageData())
    const [open, setOpen] = useState(false)
    useEffect(() => {
        messageEvent.once("message", (sub: MessageData) => {
            if (sub.use) {
                setData(sub)
                setOpen(data.use)
                setTimeout(() => {
                    setOpen(false)
                }, data.expire)
            }
        })
    })

    return (
        <Transition.Root show={open} as={Fragment}>
            <Dialog as="div" className="relative z-10" onClose={() => { }}>
                <div className="fixed inset-0 z-10 overflow-y-auto">
                    <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                        <Transition.Child
                            as={Fragment}
                            enter="ease-out duration-300"
                            enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                            enterTo="opacity-100 translate-y-0 sm:scale-100"
                            leave="ease-in duration-200"
                            leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                            leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        >
                            <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                                <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                                    <div className="sm:flex sm:items-start">
                                        <ShowIcon type={data.type} />
                                        <div className="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                                            <Dialog.Title as="h3" className="text-lg font-medium leading-6 text-gray-900">
                                                {data.title}
                                            </Dialog.Title>
                                            <div className="mt-2">
                                                <p className="text-sm text-gray-500">
                                                    {data.content}
                                                </p>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                            </Dialog.Panel>
                        </Transition.Child>
                    </div>
                </div>
            </Dialog>
        </Transition.Root>

    )
}

const ShowIcon = (s: { type: "error" | "warning" | "success" | undefined }) => {
    if (s.type === "error") {
        return (<XCircleIcon className="h-6 w-6 text-red-600" aria-hidden="true" />)
    } else if (s.type === "warning") {
        return (<ExclamationCircleIcon className="h-6 w-6 text-yellow-600" aria-hidden="true" />)
    } else if (s.type === "success") {
        return (<CheckCircleIcon className="h-6 w-6 text-green-600" aria-hidden="true" />)
    }
    return (<></>)
}

