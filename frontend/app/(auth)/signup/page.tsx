"use client";
import SignupForm from "../../components/SignupForm";
import { useAuthStore } from "@/app/store/useAuthStore";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

const Signup = () => {
    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);
    const router = useRouter();
    useEffect(() => {
        if (isLoggedIn) {
            router.replace("/");
        }
    }, [isLoggedIn, router]);

    return (
        <div className="flex flex-col w-full items-center justify-center gap-y-[16px]">
            <h1 className="text-muted-foreground text-[24px] max-[900]:text-[20px] text-center w-full">
                Create your account
            </h1>

            <SignupForm />
        </div>
    );
};

export default Signup;
