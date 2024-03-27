import {useRouter} from 'next/router'
import {useEffect, useState} from "react";
import axios from "@/utils/axios";
import {StatusPage} from "@/model/StatusPage";
import {CurrentStatus} from "@/components/CurrentStatus";
import {Separator} from "@/components/ui/separator";
import {Outages} from "@/components/Outages";
import {RecommendCompany} from "@/components/RecommendCompany";

export default function Status() {
    const [isLoading, setIsLoading] = useState(true);
    const [isError, setIsError] = useState(false)
    const [companyDisplayName, setCompanyDisplayName] = useState("");
    const [statusPageUrl, setStatusPageUrl] = useState("")
    const [isOkay, setIsOkay] = useState(false)
    const [lastStatusScrapeTime, setLastStatusScrapeTime] = useState("")
    const [incidents, setIncidents] = useState([])
    const router = useRouter()
    const companyName = router.query.company

    useEffect(() => {
        const getStatusPageInfo = async () => {
            try {
                const response = await axios.get(
                    '/api/v1/statusPage?statusPageName=' + companyName
                );
                const statusPageDetails: StatusPage = response.data.statusPage
                setCompanyDisplayName(statusPageDetails.Name)
                setStatusPageUrl(statusPageDetails.URL)
                setLastStatusScrapeTime(statusPageDetails.LastCurrentlyScraped)
            } catch (err) {
                // setIsError(true)
                console.log(err);
            }
        };
        const getCurrentStatus = async () => {
            try {
                const response = await axios.get(
                    '/api/v1/currentStatus?statusPageUrl=' + statusPageUrl
                );
                setIsOkay(response.data.isOkay)
            } catch (err) {
                // setIsError(true)
                console.log(err);
            }
        };
        const getIncidents = async () => {
            try {
                const response = await axios.get(
                    '/api/v1/incidents?statusPageUrl=' + statusPageUrl
                );
                setIncidents(response.data.incidents)
            } catch (err) {
                // setIsError(true)
                console.log(err);
            }
        };

        getStatusPageInfo().then(() => {
            getCurrentStatus().then(
                () => getIncidents().then(
                    () => setIsLoading(false)))
        });

    }, [companyName, statusPageUrl]);

    if (isLoading) {
        return <div> Loading... </div>
    }

    if (isError) {
        return <div className={"flex justify-center w-full"}>
            <div className={"sm:w-[90vw] lg:w-[80vw]"}>
                <RecommendCompany input={companyName}/>
            </div>
        </div>
    }

    return (
        <div className={"flex justify-center w-full"}>
            <div className={"sm:w-[90vw] lg:w-[80vw]"}>
                <CurrentStatus
                    displayName={companyDisplayName}
                    isOkay={isOkay}
                    lastCurrentlyScraped={lastStatusScrapeTime}
                />
                <Separator/>
                {incidents != undefined && incidents.length > 0 &&
                    <Outages incidents={incidents} statusPageUrl={statusPageUrl} displayName={companyDisplayName}/>
                }
            </div>
        </div>)
}