import Link from "next/link";
import OpenIDSection from "./OpenIDSection";

const SigninForm = () => {
    return (
        <form className="flex flex-col gap-y-[8px] justify-center border-1 border-muted-foreground p-[20px] rounded-[8px] w-full max-w-[600px] min-[901]:text-[20px]">
            <div className="flex flex-col w-full">
                <label htmlFor="email">Email</label>
                <input
                    id="email"
                    name="email"
                    type="email"
                    required
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none"
                    placeholder="your.email@example.com"
                />
            </div>

            <div className="flex flex-col w-full">
                <div className="flex justify-between">
                    <label htmlFor="password">Password</label>
                    <Link
                        href={"/forgotpassword"}
                        className="font-light text-muted-foreground cursor-pointer hover:text-foreground duration-300"
                    >
                        Forgot Password?
                    </Link>
                </div>
                <input
                    id="password"
                    name="password"
                    type="password"
                    required
                    className="border-muted-foreground border-1 rounded-[4px] p-[8px] outline-none"
                    placeholder="Enter your password"
                />
            </div>

            <button className="bg-foreground text-background rounded-[4px] py-[4px] cursor-pointer hover:bg-muted hover:text-white duration-300 border-1 border-muted-foreground">
                Sign In
            </button>

            <OpenIDSection />

            <section className="flex gap-x-[4px] w-full justify-center items-center">
                <p className="text-muted-foreground">Don't have an account? </p>
                <Link href={"/signup"} className="font-bold">
                    Sign up
                </Link>
            </section>
        </form>
    );
};

export default SigninForm;
