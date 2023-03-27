import {api} from "@/apiCalls/axiosConfig";

export const RunCode = (problemId,userId,language,timeLimit,memoryLimit,code) => {
    const body = {
        problem_Id : problemId,
        user_id : userId,
        language : language,
        time_limit : timeLimit,
        memory_limit: memoryLimit,
        SourceCode: code
    }
    api.post()
}
