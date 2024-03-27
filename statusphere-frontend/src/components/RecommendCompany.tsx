import {Card, CardHeader, CardTitle} from "@/components/ui/card";
interface RecommendCompanyProps {
    input: string;
}

export function RecommendCompany(props: RecommendCompanyProps) {
    return <div className={"flex justify-center w-full"}>
        <div className={"sm:w-[80vw] lg:w-[60vw]"}>
            <Card className={"bg-sky-400"}>
                <CardHeader className={"items-left"}>
                    <CardTitle>
                        <h3 className="scroll-m-20 text-xl font-semibold tracking-tight">
                            We currently don't have status updates for {props.input}
                        </h3>
                    </CardTitle>
                </CardHeader>
            </Card>
        </div>
    </div>
}