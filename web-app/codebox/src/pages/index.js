import Sidebar from "../components/Sidebar/Sidebar";
import Titlebar from "../components/Titlebar/Titlebar";
import QuestionCard from "../components/Question/QuestionCard";
import {useEffect, useState} from "react";
import { useRouter } from 'next/router'
import {useLocalStorage} from "../hooks/useLocalStorage";
import {v4 as uuidv4} from "uuid";
import useAxios from "axios-hooks";
import QuestionLoadingCard from "../components/Question/QuestionLoadingCard";
import {BaseUrl} from "../apiCalls/axiosConfig";
import {useHistoryLocalStorage} from "../hooks/useHistoryLocalStorage";
import Head from 'next/head'

export default function Home() {

    const [history, setHistory] = useHistoryLocalStorage('history', `{"questions":[]}`);
    const [sideBarChoice,setSideBarChoice ] = useState(0)
    const router = useRouter()
    const [id] = useLocalStorage('userId',uuidv4())
    const [display,setDisplay] = useState(false)
    const [{response = null, loading, error}, executeGetList] = useAxios({
            method: 'GET',
            url: BaseUrl+'/api/v1/question/list',

        },);
    useEffect(() =>{
        if (response !== null && error === null){
            console.log(response.data.data.Results)
            setDisplay(true)
        }
    },[response,error,loading])



  return (
    <div>
        <Head>
            <title>CodeBox</title>
            <link rel="icon" href="/codebox_logo.svg" />
        </Head>
      <main className={`bg-dullWhite h-screen`}>
          <Sidebar active={0}/>
          <Titlebar/>

          <div className={`${!display ? 'block' : 'hidden'} grid grid-cols-3`}>
              <QuestionLoadingCard />
              <QuestionLoadingCard/>
              <QuestionLoadingCard/>
          </div>

          <div className={`${display ? 'block' : 'hidden'} grid grid-cols-3 `}>
                  {display && response.data.data.Results.map((item,key) => (
                          <div className={`${display ? 'block' : 'hidden'}`} onClick={() => {
                              console.log(key)
                              console.log(response.data.data.Results[key])
                              router.push({
                                  pathname:"/problem",
                                  query :{problemData : JSON.stringify(response.data.data.Results[key]) }
                              },'/problem', { shallow: true })
                          }} key={key}>
                              <QuestionCard title={item.title} description={item.problem_statement}
                                            level={item.level} completed={false}/>
                          </div>

                      ))
                  }
          </div>




      </main>
    </div>
  )
}
