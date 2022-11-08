import {createArticle} from "../../../api/article";
import {useState} from "react";
import {Simulate} from "react-dom/test-utils";
import {MessageData, showMessage} from "../../component/message";
import input = Simulate.input;
import {Link} from "react-router-dom";

export const AddArticle = () => {
    const [data, setData] = useState({
        title: "",
        content: "",
        image: "",
    })

    const addArticle = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (data.title === "") {
            return showMessage(new MessageData(true, "error", "", "标题不能为空"))
        }
        if (data.content === "") {
            return showMessage(new MessageData(true, "error", "", "内容不能为空"))
        }
        createArticle(data).then(r => showMessage(new MessageData(true, "success", "", "新增成功")))
    }
    return (
        <>
            <div className="container mx-auto min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
                <div className="md:grid md:grid-cols-1 md:gap-6 justify-center">
                    <form method="POST" onSubmit={addArticle}>
                        <div className="shadow sm:overflow-hidden sm:rounded-md">
                            <div className="space-y-6 bg-white px-4 py-5 sm:p-6">
                                <div className="grid grid-cols-3 gap-6">
                                    <div className="col-span-3 sm:col-span-2">
                                        <label htmlFor="title"
                                               className="block text-sm font-medium text-gray-700">
                                            标题
                                        </label>
                                        <div className="mt-1 flex rounded-md shadow-sm">
                                            <input
                                                type="text"
                                                name="title"
                                                id="title"
                                                className="block w-full flex-1 rounded-none rounded-r-md border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                                placeholder="请输入标题..."
                                                onChange={e => {
                                                    data.title = e.target.value
                                                    setData(data)
                                                }
                                                }
                                            />
                                        </div>
                                    </div>
                                </div>

                                <div>
                                    <label htmlFor="content" className="block text-sm font-medium text-gray-700">
                                        内容
                                    </label>
                                    <div className="mt-1">
                                        <textarea
                                            id="content"
                                            name="content"
                                            rows={3}
                                            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                            placeholder="请输入内容..."
                                            defaultValue={''}
                                            onChange={e => {
                                                data.content = e.target.value
                                                setData(data)
                                            }
                                            }
                                        />
                                    </div>
                                </div>

                                <div>
                                    <label className="block text-sm font-medium text-gray-700">图片</label>
                                    <div
                                        className="mt-1 flex justify-center rounded-md border-2 border-dashed border-gray-300 px-6 pt-5 pb-6">
                                        <div className="space-y-1 text-center">
                                            <svg
                                                className="mx-auto h-12 w-12 text-gray-400"
                                                stroke="currentColor"
                                                fill="none"
                                                viewBox="0 0 48 48"
                                                aria-hidden="true"
                                            >
                                                <path
                                                    d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02"
                                                    strokeWidth={2}
                                                    strokeLinecap="round"
                                                    strokeLinejoin="round"
                                                />
                                            </svg>
                                            <div className="flex text-sm text-gray-600">
                                                <label
                                                    htmlFor="file-upload"
                                                    className="relative cursor-pointer rounded-md bg-white font-medium text-indigo-600 focus-within:outline-none focus-within:ring-2 focus-within:ring-indigo-500 focus-within:ring-offset-2 hover:text-indigo-500"
                                                >
                                                    <span>Upload a file</span>
                                                    <input id="file-upload" name="file-upload" type="file"
                                                           className="sr-only"/>
                                                </label>
                                                <p className="pl-1">or drag and drop</p>
                                            </div>
                                            <p className="text-xs text-gray-500">PNG, JPG, GIF up to 10MB</p>
                                        </div>
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
                                <Link
                                to="/myArticle"
                                    type="submit"
                                    className="ml-3 bg-gray-400 inline-flex justify-center rounded-md border border-transparent py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                                >
                                    返回
                                </Link>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </>
    )
}