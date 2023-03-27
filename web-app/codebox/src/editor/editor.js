import React, {useEffect, useState} from 'react';
import Editor, {ControlledEditor, monaco} from "@monaco-editor/react";
import {cppBaseCode,javaBaseCode,pythonBaseCode} from "../editor/defaultBaseCode";
import {btoa} from "buffer";

const CodeEditor = ({language, baseCode, onCodeChange,readOnly,height}) => {
    const [basecode, setBasecode] = useState('//Code Here')

    useEffect(()=>{
        if (language === 'cpp'){
            setBasecode(atob(cppBaseCode))
            onCodeChange(atob(cppBaseCode))

        }else if (language === 'java'){
            setBasecode(atob(javaBaseCode))
            onCodeChange(atob(javaBaseCode))

        }else if (language === 'python'){
            setBasecode(atob(pythonBaseCode))
            onCodeChange(atob(pythonBaseCode))
        }
    },[language])

    function handleEditorChange(value, event) {
        console.log("here is the current model value:", value);
        onCodeChange(value)
    }

    return (
        <div className="mt-5 rounded-md w-full  shadow-4xl">
            <Editor

                height={height}
                width={`100%`}
                theme={'vs-dark'}
                onChange={handleEditorChange}

                language={language || "java"}
                value={basecode}
                options={{
                    readOnly:readOnly,
                    minimap: {enabled:false,},
                    scrollBeyondLastLine:false,
                    automaticLayout:true,
                }
                }
                defaultValue={basecode}
            />
        </div>
    );
};

export default CodeEditor;