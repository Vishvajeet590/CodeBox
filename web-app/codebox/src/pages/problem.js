import React, {useEffect, useState} from 'react';
import Sidebar from "../components/Sidebar/Sidebar";
import CodeEditor from "../editor/editor";
import {v4 as uuidv4} from 'uuid';
import useAxios from "axios-hooks";
import LanguageDropdown from "../editor/menu";
import {useRouter} from "next/router";
import {useHistoryLocalStorage} from "../hooks/useHistoryLocalStorage";
import {useLocalStorage} from "../hooks/useLocalStorage";
import {BaseUrl} from "../apiCalls/axiosConfig";
import Head from "next/head";


const Problem = () => {
    const [userId] = useLocalStorage('userId', uuidv4())
    const [sourceCode, setSourceCode] = useState('')
    const [language, setLanguage] = useState('cpp')
    const [taskId, setTaskId] = useState("")
    const [polling, setPolling] = useState(true)
    const [display,setDisplay] = useState(false)
    const [running,setRunning] = useState(false)
    const [problemData,setProblemData] = useState({})
    const [history, setHistory] = useHistoryLocalStorage('history', `{"questions":[]}`);
    const router = useRouter();

    useEffect(() => {
        console.log("New querry"); // Alerts 'Someone'
        setProblemData(JSON.parse(router.query.problemData))
    }, [router.query]);


    const [{response = null, loading, error}, executePost] = useAxios({
            method: 'POST',
            url: BaseUrl+'/api/v1/runner/runCode',
            data: {  // no need to stringify
                problem_Id: problemData.question_id,
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
        if (response !== null) {
            console.log("data:", response.data.data);
            setPolling(true)
            setTaskId(response.data.data.id)
        }
    }, [response]);


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
            console.log("data:", resultResponse);
            setPolling(false)
            setRunning(false)
            if (history === null) return

            let historyObj = JSON.parse(history)
            let resultObj
            if (resultResponse.data.data.status === 'FAIL' || resultResponse.data.data.error_info === "COMPILE_ERROR"){
                resultObj = {
                    id : problemData.question_id,
                    name : problemData.title,
                    level : problemData.level,
                    status : false,
                    runtime : 0,
                    memory : 0
                }
            }else if (resultResponse.data.data.status === 'PASS'){
                resultObj = {
                    id : problemData.question_id,
                    name : problemData.title,
                    level : problemData.level,
                    status : true,
                    runtime : resultResponse.data.data.runtime_time,
                    memory : resultResponse.data.data.runtime_memory
                }
            }
            historyObj.questions.push(resultObj)
            setHistory(JSON.stringify(historyObj))
        }
    }, [resultResponse]);


    //Editor code
    const onCodeChange = (value) => {
        console.log("From Problem.js " + value)
        setSourceCode(value)
    }
    const onSubmitCode =  () => {
        setRunning(true)
        executePost().then((response) =>{
            console.log("data:", response);
        }).catch(() => {
            setRunning(true)
            console.log("ERRRRR");
        }).finally(() =>{
            setRunning(true)
        })
    };

    const onLanguageChange = (value) => {
        setLanguage(value)
        //Do editor change
        console.log("Lang = " + value)
    }

    useEffect(()=>{
        console.log("Display = "+display)
    },[display])


    return (
        <div className="w-full min-h-screen bg-dullWhite">
            <Head>
                <title>{Object.keys(problemData).length !== 0 && problemData.title}</title>
                <link rel="icon" href="/codebox_logo.svg" />
            </Head>
            <Sidebar active={-1}/>
            <div className={`grid grid-cols-2 ml-20 mr-5`}>
                <div className={`flex flex-col mr-2`}>
                    <div className={`flex flex-row gap-6 items-center mt-5`}>
                        <p className={`text-black font-nunito font-semibold text-2xl`}>{Object.keys(problemData).length !== 0 && problemData.title}</p>
                        <p className={`text-AcceptedStrip font-nunito font-bold text-lg `}>Easy</p>
                    </div>
                    <p className={`text-black font-nunito font-semibold text-xl mt-5`}>Problem</p>
                    <p className={`text-black font-nunito font-medium text-md text-gray-900 mt-2`}>
                        {Object.keys(problemData).length !== 0 && problemData.problem_statement}
                    </p>

                    <p className={`text-black font-nunito font-semibold text-xl mt-5`}>Input</p>
                    <p className={`text-black font-nunito font-medium text-md text-gray-900 mt-2`}>
                        {Object.keys(problemData).length !== 0 && problemData.input_statement}

                    </p>

                    <p className={`text-black font-nunito font-semibold text-xl mt-5`}>Output</p>
                    <p className={`text-black font-nunito font-medium text-md text-gray-900 mt-2`}>
                        {Object.keys(problemData).length !== 0 && problemData.output_statement}

                    </p>

                    <p className={`text-black font-nunito font-semibold text-xl mt-5`}>Sample 1:</p>

                    <div className={`relative overflow-hidden mb-10`}>

                        <table className="min-w-full">
                            <thead className={`border-b border-b-gray-400`}>
                            <tr>
                                <th className={`text-sm font-bold font-nunito text-gray-900 px-6 py-4 text-left`}>Input</th>
                                <th className={`text-sm font-bold font-nunito text-gray-900 px-6 py-4 text-left`}>Output</th>
                            </tr>
                            </thead>
                            <tbody>
                            {Object.keys(problemData).length !== 0 && problemData.sample_inputs.map((item,key)=>(
                                <tr key={key}>
                                    <td className={`px-6 py-1 whitespace-nowrap text-sm font-medium text-gray-900`}>
                                        <pre>{item}</pre>
                                    </td>
                                    <td className={`px-6 py-1 whitespace-nowrap text-sm font-medium text-gray-900`}>
                                        <pre>{problemData.sample_outputs[key]}</pre>
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>


                </div>
                <div className={`flex flex-col`}>
                    <LanguageDropdown onLanguageChange={onLanguageChange}/>
                    <div className={`flex flex-row justify-end mt-5`}>
                        <button onClick={onSubmitCode} type="button"
                                disabled={running}
                                className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium font-nunito rounded-lg text-sm px-5 py-2 mr-2  focus:outline-none">
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
                    <CodeEditor height={'60vh'} onCodeChange={onCodeChange} baseCode={`//Code here`} language={language}
                                readOnly={false}/>

                    <div>
                        <p className={`font-nunito text-lg text-gray-900 underline underline-offset-2 p-1 mt-5`}>Result</p>
                        <div className={`flex flex-col gap-4 mt-2 ${display && resultResponse.data.data.status === "PASS" ?'block' : 'hidden'}`}>
                            <p className={`font-nunito text-Accepted text-3xl `}>Accepted</p>

                            <div className={`rounded-lg w-full flex flex-row gap-5 bg-Accepted bg-opacity-10 p-5 mr-5 `}>
                                <div className={`flex flex-row items-center gap-2`}>
                                    <p className={`font-nunito text-Accepted text-lg font-bold`}>Runtime </p>
                                    <p className={`font-nunito text-Accepted text-lg`}>{display && resultResponse.data.data.runtime_time/10} ms </p>
                                </div>
                                <div className={`flex flex-row items-center gap-2`}>
                                    <p className={`font-nunito text-Accepted text-lg font-bold`}>Memory </p>
                                    <p className={`font-nunito text-Accepted text-lg`}>{display && resultResponse.data.data.runtime_memory} kb</p>
                                </div>
                            </div>
                        </div>

                        <div
                            className={`flex flex-col gap-4 ${display && resultResponse.data.data.error_info === "COMPILE_ERROR" ?'block' : 'hidden'}`}>
                            <p className={`font-nunito text-Rejected text-2xl font-semibold`}>Compile Error</p>
                            <div className={`rounded-lg w-full bg-Rejected bg-opacity-10 p-5 mr-5`}>
                                <p className={`font-nunito text-Rejected`}>
                                    {display && atob(resultResponse.data.data.wrong_line)}
                                </p>
                            </div>
                        </div>

                        <div
                            className={`flex flex-col gap-4 ${display && resultResponse.data.data.error_info !== '' && resultResponse.data.data.error_info !== 'COMPILE_ERROR'  ?'block' : 'hidden'}`}>
                            <p className={`font-nunito text-Rejected text-2xl font-semibold`}>Compile Error</p>
                            <div className={`rounded-lg w-full bg-Rejected bg-opacity-10 p-5 mr-5`}>
                                <p className={`font-nunito text-Rejected`}>
                                    {display && 'Python compilation error.'}
                                </p>
                            </div>
                        </div>
                        <div
                            className={`flex flex-col mb-10 gap-4 ${display && resultResponse.data.data.status === "FAIL"  ?'block' : 'hidden'}`}>
                            <p className={`font-nunito text-Rejected text-2xl font-semibold`}>Rejected</p>
                            <div className={`rounded-lg w-full bg-Rejected bg-opacity-10 p-5 mr-5`}>
                                <p className={`font-nunito text-Rejected`}>
                                    Some testcases failed.
                                </p>

                                <div className={`flex flex-row gap-10 mt-3`}>
                                    <p className={`font-nunito font-bold`}>Expected</p>
                                    <pre className={`font-nunito font-bold`}>Output</pre>
                                </div>
                                <div className={`flex flex-row gap-20`}>
                                    <pre className={`ml-2 text-right`}>{display && resultResponse.data.data.expected_output}</pre>
                                    <pre className={`ml-2 text-left`}>{display && resultResponse.data.data.last_output}</pre>
                                </div>
                            </div>
                        </div>
                        <div
                            className={`flex flex-col gap-4 ${''}`}>

                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Problem;