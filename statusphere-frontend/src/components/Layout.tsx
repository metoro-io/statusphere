import {Header} from "@/components/header";
import {Footer} from "@/components/footer";
import {Meteors} from "@/components/mainsite/Meteors";

const Layout = ({children}) => {
    return (
        <>
            <Header/>
            <Meteors number={20}/>
            <div className={"z-10"}>
                {children}
            </div>
            <Footer/>
        </>
    );
};
export default Layout;