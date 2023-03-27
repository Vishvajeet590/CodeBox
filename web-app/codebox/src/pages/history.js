import React, {useEffect, useState} from 'react';
import {RiHistoryLine} from 'react-icons/ri';
import Sidebar from "../components/Sidebar/Sidebar";
import {useHistoryLocalStorage} from "../hooks/useHistoryLocalStorage";
import emptyLog from "../../public/empty.svg"
import Image from 'next/image'
import Head from "next/head";


const History = () => {

    const [history, setHistory] = useHistoryLocalStorage('history', `{"questions":[]}`);
    const [historyObj, setHistoryObj] = useState(null)
    const [display, setDisplay] = useState(false)

    useEffect(() => {
        console.log("history", history)
        let obj = JSON.parse(history)
        setHistoryObj(obj)
    }, [history])

    useEffect(() => {
        if (historyObj !== null && historyObj.questions.length > 0 ){
            console.log(historyObj)
            setDisplay(true)
        }else {
            console.log("No data")
        }
    },[historyObj])


    return (
        <div className={`w-full min-h-screen bg-dullWhite`}>
            <Head>
                <title>History</title>
                <link rel="icon" href="/codebox_logo.svg" />
            </Head>
            <Sidebar active={2}/>
            <div className={`flex flex-row ml-20 items-center gap-3`}>
                <h1 className={`text-xl font-bold font-nunito mt-5`}>Your History</h1>
                <RiHistoryLine className={`text-2xl text-gray-500 mt-5`}/>
            </div>

            <div className={`overflow-auto rounded-lg shadow hidden md:block mx-20 my-8`}>
                <table className={`w-full`}>
                    <thead className={`bg-gray-50 border-b-2 border-gray-200`}>
                    <tr>
                        <th className="w-20 p-3 text-sm font-semibold tracking-wide text-left">No.</th>
                        <th className="w-20 p-3 text-sm font-semibold tracking-wide text-left">Problem</th>
                        <th className="w-20 p-3 text-sm font-semibold tracking-wide text-left">level</th>
                        <th className="w-20 p-3 text-sm font-semibold tracking-wide text-left">Runtime</th>
                        <th className="w-20 p-3 text-sm font-semibold tracking-wide text-left">Memory</th>
                        <th className="w-20 p-3 text-sm font-semibold tracking-wide text-left">Status</th>
                    </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-100">
                    {
                        display &&  historyObj.questions.map((item,key) => (
                            <tr key={key} className="odd:bg-white even:bg-slate-50">
                                <td className="p-3 text-sm text-gray-700 whitespace-nowrap">
                                    <a className="font-bold text-blue-500">{key+1}</a>
                                </td>
                                <td className="p-3 text-sm text-gray-700 whitespace-nowrap">
                                    {item.name}
                                </td>
                                <td className="p-3 text-sm text-gray-700 whitespace-nowrap">{item.level}</td>
                                <td className="p-3 text-sm text-gray-700 whitespace-nowrap">{item.runtime} ms</td>
                                <td className="p-3 text-sm text-gray-700 whitespace-nowrap">{item.memory} kb</td>
                                <td className="p-3 text-sm text-gray-700 whitespace-nowrap"><span className={`p-2 text-xs font-medium uppercase tracking-wider ${item.status === true ? 'text-green-800 bg-green-200' : 'text-red-800 bg-red-200'}  rounded-lg bg-opacity-50`}>{item.status === true ?'PASS' : 'FAIL'}</span></td>
                            </tr>
                        ))
                    }
                    </tbody>
                </table>
            </div>
            <div className={`flex flex-row justify-center ${display ? 'hidden' : 'block'}`}>
                <Image width={300} height={300} src={emptyLog} alt={"empty logo"}/>
            </div>

            <div className={`flex flex-row justify-center ${display ? 'hidden' : 'block'}`}>
                <p className={`font-nunito font-semibold text-darkBlue mt-5`}>It's Empty here... Try out some problems</p>
            </div>
        </div>
    );
};

export default History;