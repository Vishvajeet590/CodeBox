import React, {useEffect, useMemo, useState} from 'react';
import { createAvatar } from '@dicebear/core';
import { bottts } from '@dicebear/collection';
import {useLocalStorage} from "../../hooks/useLocalStorage";
import {v4 as uuidv4} from "uuid";



const Titlebar = () => {

    const [avatar,setAvatar] = useState('')

    const [seed] = useLocalStorage('seed',uuidv4())
    useEffect(()=>{

        const av = createAvatar(bottts, {
            seed: seed,
            backgroundType: [
                "gradientLinear"
            ],
            backgroundColor: [
                "ffdfbf",
                "ffd5dc",
                "d1d4f9",
                "c0aede",
                "b6e3f4"
            ],
            randomizeIds: true,
            dataUri: true,
            size: 52,
            radius: 20
        }).toDataUriSync();
        setAvatar(av)
    },[])


    //TODO Create that yesterday tomorrow question thing, later
    return (
        <div className={`w-full flex flex-row justify-between`}>
            <div className={`mx-24 flex flex-row gap-5 mt-5`}>
{/*
                <p className={`font-nunito text-gray-500 font-semibold text-3xl hover:text-4xl transition-all duration-100 hover:cursor-pointer hover:text-purple ease-linear`}>Jan 21</p>
*/}
                <div className={`flex flex-col items-center`}>
                    <p className={`font-nunito text-darkBlue text-4xl`}>Today's Questions</p>
                    <h className={`h-1 w-1/3 bg-purple`}></h>
                </div>
{/*
                <p className={`font-nunito text-gray-500 font-semibold text-3xl hover:text-4xl transition-all duration-100 hover:cursor-pointer hover:text-purple ease-linear`}>Jan 23</p>
*/}
            </div>

                <img className={`mx-12 mt-5`} src={avatar} alt={"avatar"}/>
        </div>
    );
};

export default Titlebar;