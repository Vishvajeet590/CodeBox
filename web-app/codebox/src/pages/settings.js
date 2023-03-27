import React from 'react';
import Sidebar from "../components/Sidebar/Sidebar";
import progress from "../../public/progress.svg"
import Image from "next/image";
import Head from "next/head";
const Settings = () => {
    return (
        <div className={`w-full min-h-screen bg-dullWhite flex flex-col justify-center items-center`}>
            <Head>
                <title>CodeBox</title>
                <link rel="icon" href="/codebox_logo.svg" />
            </Head>
            <Sidebar active={3}/>
            <div className={`flex flex-row justify-center`}>
                <Image width={400} height={400} src={progress} alt={"progress"}/>
            </div>
            <div className={`flex flex-row justify-center`}>
                <p className={`font-nunito font-semibold text-darkBlue mt-5`}>Working on this part.</p>
            </div>


        </div>
    );
};

export default Settings;