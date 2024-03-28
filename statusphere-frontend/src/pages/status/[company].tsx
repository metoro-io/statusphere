import {useRouter} from 'next/router'
import {useEffect, useState} from "react";
import axios from "@/utils/axios";
import {StatusPage} from "@/model/StatusPage";
import {Status} from "@/model/Status";
import {CurrentStatus} from "@/components/CurrentStatus";
import {Separator} from "@/components/ui/separator";
import {Outages} from "@/components/Outages";
import {RecommendCompany} from "@/components/RecommendCompany";

export default function CompanyStatusPage() {
    const [isLoading, setIsLoading] = useState(true);
    const [isError, setIsError] = useState(false)
    const [currStatus, setCurrStatus] = useState(Status.UNKNOWN)
    const [statusPageDetails, setStatusPageDetails] = useState<StatusPage>({} as StatusPage);
    const [companyName, setCompanyName] = useState<string>("");
    const router = useRouter()

    useEffect(() => {
        const getStatusPageInfo = async () => {
            try {
                const statusPageResp = await axios.get(
                    '/api/v1/statusPage?statusPageName=' + router.query.company
                );
                const statusPageDetails: StatusPage = statusPageResp.data.statusPage
                setStatusPageDetails(statusPageDetails)
                const currStatusResp = await axios.get(
                    '/api/v1/currentStatus?statusPageUrl=' + statusPageDetails.url
                );
                setCurrStatus(currStatusResp.data.status)
            } catch (err) {
                setIsError(true)
                console.log(err);
            }
        };
        if (router.query.company != undefined) {
            setCompanyName(router.query.company as string)
            getStatusPageInfo().then(() => setIsLoading(false));
        }
    }, [router.query.company]);

    if (isLoading) {
        return <div> Loading... </div>
    }

    return (
        <div className={"flex justify-center w-full"}>
            <div className={"sm:w-[90vw] lg:w-[80vw] z-10 space-y-8"}>
                {isError ?
                    <RecommendCompany input={companyName}/>
                    :
                    <>
                        <CurrentStatus
                            displayName={statusPageDetails.name}
                            status={currStatus}
                            lastCurrentlyScraped={statusPageDetails.lastCurrentlyScraped}
                            statusPageUrl={statusPageDetails.url}
                        />
                        <Outages statusPageDetails={statusPageDetails}/>
                    </>
                }
            </div>
        </div>)
}