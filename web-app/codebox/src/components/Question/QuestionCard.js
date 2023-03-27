import React from 'react';
import { BsCheckLg } from 'react-icons/bs';
import { MdKeyboardArrowRight } from 'react-icons/md';

const QuestionCard = ({title,description,level,completed}) => {
    return (
        <div className={`bg-white hover:bg-gray-50 hover:cursor-pointer rounded-2xl w-96 h-36 mx-20 flex flex-row my-5`}>
            <div className={`flex flex-row items-center`}>
                <hr className={`w-1 h-3/4 mx-4 ${ completed ? `bg-AcceptedStrip` : `bg-unAttemptedStrip`}`}/>
            </div>
            <div className={`${completed ? `block` : `hidden`} drop-shadow-lg my-4 p-3 mr-5 w-10 h-10 rounded-full flex justify-center items-center bg-AcceptedStrip bg-opacity-30`}>
                <BsCheckLg className={`text-AcceptedStrip`}/>
            </div>
            <div className={`flex flex-col gap-1`}>
                <div className={`flex flex-row justify-between`}>
                    <p className={`text-black font-nunito font-bold mt-5 text-md`}>{title}</p>
                    <MdKeyboardArrowRight className={`text-gray-500 text-xl mt-4 mr-4`}/>
                </div>

                <p className={`text-gray-600 font-nunito font-semibold line-clamp-2`}>{description}</p>

                <p className={` ${level == "Easy" ? `text-AcceptedStrip` : level == "Medium" ? `text-AttemptedStrip` : `text-hard`}  font-semibold font-nunito`}>{level}</p>
            </div>
        </div>
    );
};

export default QuestionCard;