import axios from "@/utils/axios";
import {StatusPage} from "@/model/StatusPage";
import {Status} from "@/model/Status";
import {CurrentStatus} from "@/components/CurrentStatus";
import {Outages} from "@/components/Outages";
import {RecommendCompany} from "@/components/RecommendCompany";
import Head from "next/head";
import React from "react";
import {
    Drawer,
    DrawerContent,
    DrawerDescription,
    DrawerHeader,
    DrawerTitle,
    DrawerTrigger
} from "@/components/ui/drawer";
import {CopyBlock, dracula} from "react-code-blocks";

interface CompanyStatusPageProps {
    statusPageDetails: StatusPage
    currStatus: Status
    companyName: string
    outages: Incident[]
    isError?: boolean
    apiCallCurrentStatus: string
    apiCallIncidents: string
}

export async function getServerSideProps(context: any) {
    const companyName = context.params.company
    try {
        const statusPageResp = await axios.get(
            '/api/v1/statusPage?statusPageName=' + companyName
        );
        const statusPageDetails: StatusPage = statusPageResp.data.statusPage
        const currStatusResp = axios.get(
            '/api/v1/currentStatus?statusPageUrl=' + statusPageDetails.url
        );
        const outagesResp = axios.get(
            '/api/v1/incidents?limit=50&statusPageUrl=' + statusPageDetails.url
        );

        const currStatus: Status = (await currStatusResp).data.status
        const outages = (await outagesResp).data.incidents

        const apiCallCurrentStatus = `curl -X GET "https://statusphere.metoro.io/api/v1/currentStatus?statusPageUrl=${statusPageDetails.url}"`
        const apiCallIncidents = `curl -X GET "https://statusphere.metoro.io/api/v1/incidents?statusPageUrl=${statusPageDetails.url}"`

        return {
            props: {
                statusPageDetails: statusPageDetails,
                currStatus: currStatus,
                companyName: companyName,
                outages: outages,
                apiCallCurrentStatus: apiCallCurrentStatus,
                apiCallIncidents: apiCallIncidents
            }
        }
    } catch (e) {
        return {
            props: {
                isError: true,
                companyName: companyName
            }
        }
    }
}


export default function CompanyStatusPage({
                                              statusPageDetails,
                                              currStatus,
                                              companyName,
                                              outages,
                                              isError,
                                              apiCallCurrentStatus,
                                              apiCallIncidents
                                          }: CompanyStatusPageProps) {

    return (
        <div>
            <div className={"flex justify-center px-[10vw]"}>
                <div className={"flex justify-start w-full"}>
                    <h1 className={"text-2xl"}>
                        {companyName} Status
                    </h1>
                </div>
                <div className={"flex justify-end w-full"}>
                    <Drawer>
                        <DrawerTrigger asChild>
                            <button
                                className={"text-sm text-gray-500"}>
                                Show API Calls
                            </button>
                        </DrawerTrigger>
                        <DrawerContent className={"bg-white"}>
                            <div className="mx-auto w-full max-w-xl">
                                <DrawerHeader>
                                    <DrawerTitle className={"text-2xl"}>API calls</DrawerTitle>
                                    <DrawerDescription>Programmatically
                                        get {statusPageDetails.name} status.</DrawerDescription>
                                </DrawerHeader>
                                {/*The current status api call*/}
                                <div>
                                    <div className={"text-lg font-semibold mb-4"}> Get current status</div>
                                    <CopyBlock language={"bash"} text={apiCallCurrentStatus} showLineNumbers={false}
                                               wrapLongLines={true} theme={dracula}/>
                                </div>
                                {/*The incidents api call*/}
                                <div className={"mb-4"}>
                                    <div className={"text-lg font-semibold mb-4 mt-4"}> Get historical incidents</div>
                                    <CopyBlock language={"bash"} text={apiCallIncidents} showLineNumbers={false}
                                               wrapLongLines={true} theme={dracula}/>
                                </div>
                            </div>
                        </DrawerContent>
                    </Drawer>
                </div>
            </div>
            <div className={"mt-4 flex justify-center w-full z-10"}>
                <Head>
                    <title>{statusPageDetails.name} Status - Statusphere</title>
                    <meta name="description"
                          content={`Current status of ${statusPageDetails.name}. Is ${statusPageDetails.name} down?`}/>
                    <meta name="keywords"
                          content={`status, statusphere, statuspage, up, down, ${statusPageDetails.name}`}/>
                </Head>


                <div className={"w-[90vw] lg:w-[80vw] space-y-8 flex justify-center"}>
                    <div>
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
                                <Outages statusPageDetails={statusPageDetails} incidents={outages}/>
                            </>
                        }
                    </div>
                </div>
            </div>
        </div>)
}