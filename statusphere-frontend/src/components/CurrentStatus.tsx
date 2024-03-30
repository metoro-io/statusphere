import {Card, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Status} from "@/model/Status";
import {timeAgo} from "@/utils/datetime";

interface CurrentStatusProps {
    displayName: string;
    status: string;
    lastCurrentlyScraped: string;
    statusPageUrl: string;
}

export function CurrentStatus(props: CurrentStatusProps) {

    if (props.status === Status.UNKNOWN) {
        return <div></div>
    }

    return <Card className={getStatusColor(props.status)}>
        <CardHeader className={"items-left"}>
            <CardTitle className={"scroll-m-20 text-xl font-semibold tracking-tight"}>
                Is {props.displayName} down?
            </CardTitle>
            <CardDescription className={"scroll-m-20 border-b pb-2 text-2xl font-semibold tracking-tight first:mt-0"}>
                    Current Status: {props.displayName} is {props.status}
            </CardDescription>
        </CardHeader>
        <CardFooter>
            <p className="leading-7 [&:not(:first-child)]:mt-6">
                We checked the <a href={props.statusPageUrl}> official {props.displayName} status page </a>
                for updates {timeAgo(props.lastCurrentlyScraped)}.
            </p>
        </CardFooter>
    </Card>
}

const getStatusColor = (status: string): string => {
    if (status.toUpperCase() === Status.UP) {
        return "bg-green-300"
    }
    if (status.toUpperCase() === Status.DOWN || status.toUpperCase() === Status.DEGRADED) {
        return "bg-red-300"
    }
    return "bg-blue-300";
}
