import {BsPlus, BsFillLightningFill, BsGearFill} from 'react-icons/bs';
import {FaFire, FaPoo} from 'react-icons/fa';
import {BsJournalCode,BsGithub} from 'react-icons/bs';
import {RiHistoryLine} from 'react-icons/ri';
import Image from "next/image";
import Logo from "../../../public/codebox_logo.svg"
import {useRouter} from "next/router";

const SideBar = ({active}) => {
    const router = useRouter()



    const onButtonClick = (e) => {
        if (e.target.id === 'problem' && active != 0) {
            //Route to problems only if on other pages
            console.log("Moving to problems...")
            router.push("/")
        } else if (e.target.id === 'run' && active != 1) {
            //Route to problems only if on other pages
            console.log("Moving to Run...")
            router.push("/sandbox")
        }else if (e.target.id === 'history' && active != 2){
            //Route to problems only if on other pages
            console.log("Moving to history...")
            router.push("/history")

        }else if (e.target.id === 'setting' && active != 3){
            //Route to problems only if on other pages
            console.log("Moving to setting...")
            router.push("/settings")
        }
    }


    const onGithubCLick = (e)=>{
        const url = 'https://github.com/Vishvajeet590';
        window.open(url, '_blank')
    }

    return (
        <div className="z-10 fixed top-0 left-0 h-screen w-16 flex flex-col bg-white shadow-lg">
            {/*TODO change LOGO url*/}
            <a className={`p-1 mt-3`} href={`/`}>
                <Image src={Logo} alt={"logo"}/>
            </a>

            {/*TODO Fix the button click issue*/}
            <div className={` h-full flex flex-col justify-center gap-4`}>
                <button onClick={onButtonClick} id={'problem'}
                     className={`${active == 0 ? 'rounded-xl bg-purple text-white' : ''} hover:bg-purple text-gray-500 hover:text-white sidebar-icon group`}>
                    <BsJournalCode onClick={onButtonClick}  id={'problem'} size="28"/>
                    <span className="sidebar-tooltip group-hover:scale-100">Problem</span>
                </button>
                <button onClick={onButtonClick} id={'run'}
                     className={`${active == 1 ? 'rounded-xl bg-purple text-white' : ''} hover:bg-purple text-gray-500 hover:text-white sidebar-icon group`}>
                    <BsFillLightningFill  onClick={onButtonClick} id={'run'} size="20"/>
                    <span className="sidebar-tooltip group-hover:scale-100">Quick Snippet Run</span>
                </button>
                <div onClick={onButtonClick} id={'history'}
                     className={`${active == 2 ? 'rounded-xl bg-purple text-white' : ''} hover:bg-purple text-gray-500 hover:text-white sidebar-icon group`}>
                    <RiHistoryLine onClick={onButtonClick} id={'history'} size="24"/>
                    <span className="sidebar-tooltip group-hover:scale-100">History</span>
                </div>
                <div onClick={onButtonClick} id={'setting'}
                     className={`${active == 3 ? 'rounded-xl bg-purple text-white' : ''} hover:bg-purple text-gray-500 hover:text-white sidebar-icon group`}>
                    <BsGearFill onClick={onButtonClick} id={'setting'} size="22"/>
                    <span className="sidebar-tooltip group-hover:scale-100">Setting</span>
                </div>

            </div>

            <div onClick={onGithubCLick} id={'Github'}
                 className={`${active == 4 ? 'rounded-xl bg-purple text-white' : ''} mb-5 hover:bg-purple text-gray-500 hover:text-white sidebar-icon group`}>
                <BsGithub onClick={onGithubCLick} id={'Github'} size="28"/>
                <span className="sidebar-tooltip group-hover:scale-100">Github</span>
            </div>
        </div>
    );
};

export default SideBar;