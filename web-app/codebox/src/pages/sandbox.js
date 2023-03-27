import React, {useEffect, useState} from 'react';
import Sidebar from "../components/Sidebar/Sidebar";
import CodeEditor from "../editor/editor";
import LanguageDropdown from "../editor/menu";
import useAxios from "axios-hooks";
import {useLocalStorage} from "../hooks/useLocalStorage";
import {v4 as uuidv4} from "uuid";
import {BaseUrl} from "../apiCalls/axiosConfig";
import Head from "next/head";


const Sandbox = () => {
    const [userId] = useLocalStorage('userId', uuidv4())
    const [language, setLanguage] = useState('cpp')
    const [running,setRunning] = useState(false)
    const [sourceCode, setSourceCode] = useState('')
    const [display,setDisplay] = useState(false)
    const [taskId, setTaskId] = useState("")
    const [polling, setPolling] = useState(true)


    const [{response = null, loading, error}, executePost] = useAxios({
            method: 'POST',
            url: BaseUrl+'/api/v1/runner/runCode',
            data: {  // no need to stringify
                problem_Id: 'sandbox-101',
                user_id: userId,
                language: language.toLocaleUpperCase(),
                time_limit: 10,
                memory_limit: 20,
                SourceCode: btoa(sourceCode)
            },
        },
        {manual: true});

    const [{response: resultResponse = null, resultLoading, resultError}, executeResult] = useAxios({
            method: 'GET',
            url: BaseUrl+'/api/v1/runner/result/' + taskId,

        },
        {manual: true});

    useEffect(() => {
        console.log(taskId)
        let counter = 0;
        if (taskId !== '') {
            const interval = setInterval(() => {
                if (counter < 10 && polling) {
                    executeResult().then(() => {
                        console.log('clearing interval from then')
                        clearInterval(interval)
                        //set result object
                        setDisplay(true)
                        setRunning(false)

                    }).catch((err) => {
                        console.log(err)
                        setRunning(false)
                    })
                    counter++
                } else {
                    console.log('clearing interval')
                    setRunning(false)
                    clearInterval(interval)
                }
            }, 5000)
        }else {
            console.log("passing")
            setRunning(false)
        }
    }, [taskId])

    useEffect(() => {
        if (resultResponse !== null) {
            console.log("data:", resultResponse.data.data);
            setPolling(false)
            setRunning(false)
        }
    }, [resultResponse]);
    useEffect(() => {
        if (response !== null) {
            console.log("data:", response.data.data);
            setPolling(true)
            setTaskId(response.data.data.id)
        }
    }, [response]);

    const onSubmitCode = async () => {
        /*sendCode(true);*/
        executePost()
        setRunning(true)
        console.log("data:", response);
    };

    const onLanguageChange = (value) => {
        setLanguage(value)
        //Do editor change
        console.log("Lang = " + value)
    }

    useEffect(()=>{
        console.log("Display = "+display)
    },[display])

    const onCodeChange = (value) => {
        console.log("From Problem.js " + value)
        setSourceCode(value)
    }


    return (
        <div className="w-full min-h-screen bg-dullWhite">
            <Head>
                <title>SandBox</title>
                <link rel="icon" href="/codebox_logo.svg" />
            </Head>
            <Sidebar active={1}/>
            <div className={`grid grid-cols-2`}>
                <div className={`ml-20`}>
                    <LanguageDropdown onLanguageChange={onLanguageChange}/>
                    <div className={`flex flex-row justify-end mt-5`}>
                        <button onClick={onSubmitCode} type="button"
                                disabled={running}
                                className={`${!running ? 'disabled:opacity-70' : ''} text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium font-nunito rounded-lg text-sm px-5 py-2 mr-2  focus:outline-none`}>
                            <div className={`flex flex-row gap-2 `}>
                                {!running ? 'Submit' : 'Running'}
                                <svg aria-hidden="true" role="status"
                                     className={`${running ? 'block' : 'hidden'} inline w-4 h-4 text-white animate-spin`} viewBox="0 0 100 101"
                                     fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path
                                        d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                                        fill="#E5E7EB"/>
                                    <path
                                        d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                                        fill="currentColor"/>
                                </svg>
                            </div>
                        </button>
                    </div>
                    <CodeEditor height={'70vh'} onCodeChange={onCodeChange} baseCode={`//Code here`} language={language}
                                readOnly={false}/>
                    <div className="p-4 mt-5 mb-4 text-sm text-red-800 rounded-lg bg-yellow-50">
                        <span className="font-medium">Runtime Input</span> will be added later... Maybe !
                    </div>
                </div>

                <div className={`mx-5 mt-16`}>
                    <div className={`flex flex-col gap-4 ${display && resultResponse.data.data.error_info !== '' ? 'block' : 'hidden'  }`}>
                        <p className={`font-nunito text-Rejected text-2xl font-semibold`}>Compile Error</p>
                        <div className={`rounded-lg w-full bg-Rejected bg-opacity-10 p-5 mr-5`}>
                            <p className={`font-nunito text-Rejected`}>
                                {display && resultResponse.data.data.error_info}
                            </p>
                        </div>
                    </div>

                    <div className={`flex flex-col gap-4 mt-2 ${display && resultResponse.data.data.error_info === '' &&  resultResponse.data.data.runtime_memory > 0? 'block' : 'hidden'  }`}>
                        <p className={`font-nunito text-Accepted text-3xl `}>Output</p>

                        <div className={`rounded-lg w-full  gap-5 bg-Accepted bg-opacity-10 p-5 mr-5 `}>
                            <div className={`flex flex-row items-center gap-2`}>
                                <pre className={`font-nunito text-Accepted text-lg`}>{display && resultResponse.data.data.last_output}</pre>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    );
};

export default Sandbox;