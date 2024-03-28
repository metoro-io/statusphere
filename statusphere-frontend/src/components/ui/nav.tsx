"use client"

import {LucideIcon} from "lucide-react"
import {Tooltip, TooltipContent, TooltipTrigger} from "./tooltip"
import {cn} from "./lib/utils"
import {buttonVariants} from "./button"
import {Link} from "react-router-dom"

export interface NavLink {
    title: string
    label?: string
    icon: LucideIcon
    variant: "default" | "ghost"
    location: string
    onClick?: () => void
}

export interface NavProps {
    isCollapsed: boolean
    links: NavLink[]
}

export function Nav({links, isCollapsed}: NavProps) {
    return (
        <div
            data-collapsed={isCollapsed}
            className="group h-max flex flex-col gap-4 py-2 data-[collapsed=true]:py-2"
        >
            <nav
                className="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2 h-9 text-lg h-max">
                {links.map((link, index) =>
                        isCollapsed ? (
                            <Tooltip key={index} delayDuration={0}>
                                <TooltipTrigger asChild>
                                    <Link
                                        onClick={link.onClick}
                                        to={link.location}
                                        className={cn(
                                            buttonVariants({variant: link.variant, size: "icon"}),
                                            "h-9 w-9",
                                            link.variant === "default" &&
                                            "dark:bg-muted dark:text-muted-foreground dark:hover:bg-muted dark:hover:text-white"
                                        )}
                                    >
                                        <link.icon className="h-4 w-4"/>
                                        <span className="sr-only">{link.title}</span>
                                    </Link>
                                </TooltipTrigger>
                                <TooltipContent side="right" className="flex items-center gap-4">
                                    {link.title}
                                    {link.label && (
                                        <span className="ml-auto text-muted-foreground">
                    {link.label}
                  </span>
                                    )}
                                </TooltipContent>
                            </Tooltip>
                        ) : (
                            <Link
                                key={index}
                                to={link.location}
                                onClick={link.onClick}
                                className={cn(
                                    buttonVariants({variant: link.variant, size: "sm"}),
                                    link.variant === "default" &&
                                    "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
                                    "justify-start text-sm h-9"
                                )}
                            >
                                <link.icon className="mr-2 h-4 w-4"/>
                                {link.title}
                                {link.label && (
                                    <span
                                        className={cn(
                                            "ml-auto",
                                            link.variant === "default" &&
                                            "text-background dark:text-white"
                                        )}
                                    >
                  {link.label}
                </span>
                                )}
                            </Link>
                        )
                )}
            </nav>
        </div>
    )
}