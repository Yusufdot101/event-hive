"use client";
import { useAuthStore } from "@/app/store/useAuthStore";
import SignupForm from "../../components/SignupForm";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

const page = () => {
    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);
    const router = useRouter();
    useEffect(() => {
        if (isLoggedIn) {
            router.replace("/");
        }
    }, []);

    return (
        <div className="flex flex-col w-full items-center justify-center gap-y-[16px]">
            <h1 className="text-muted-foreground text-[24px] max-[900]:text-[20px] text-center w-full">
                Create your account
            </h1>

            <SignupForm />
        </div>
    );
};

export default page;
