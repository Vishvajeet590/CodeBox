import { Menu,Transition } from '@headlessui/react'
import {useState} from "react";

function LanguageDropdown({onLanguageChange}) {
    const [selected,setSelected] = useState('CPP')
    const onClick = ()=>{
        onLanguageChange(selected)
    }
    return (
        <Menu >
            <Menu.Button className={`absolute mt-5 bg-gray-300 pl-3 pr-1 pt-1 pb-1 hover:bg-gray-400 w-24 rounded-lg text-left `}>{selected}</Menu.Button>

            <Menu.Items className={`z-10 fixed flex flex-col bg-gray-300 w-36 rounded-lg mt-16 p-1`}>
                <Menu.Item onClick={()=>{
                    setSelected("cpp")
                    onLanguageChange('cpp')
                }} >

                    {({ active }) => (
                        <a className={`${active || selected ==='cpp' ? 'bg-blue-500 w-full rounded-md px-2 my-1' :'px-2 my-1'}`}>
                            CPP
                        </a>
                    )}
                </Menu.Item>
                <Menu.Item  onClick={()=>{
                    setSelected("java")
                    onLanguageChange('java')
                }} >
                    {({ active }) => (
                        <a className={`${active || selected ==='java' ? 'bg-blue-500 w-full rounded-md px-2 my-1' :'px-2 my-1'}`}>
                            Java
                        </a>
                    )}
                </Menu.Item>
                <Menu.Item onClick={()=>{
                    setSelected("python")
                    onLanguageChange('python')
                }}>
                    {({ active }) => (
                        <a className={`${active || selected ==='python' ? 'bg-blue-500 w-full rounded-md px-2 my-1' :'px-2 my-1'}`}>
                            Python
                        </a>
                    )}
                </Menu.Item>

            </Menu.Items>
        </Menu>
    )
}
export default LanguageDropdown
