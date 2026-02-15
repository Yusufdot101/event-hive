import SigninForm from "../../components/SigninForm";

const page = () => {
    return (
        <div className="flex flex-col w-full items-center justify-center gap-y-[16px]">
            <h1 className="text-muted-foreground text-[24px] max-[900]:text-[20px] text-center w-full">
                Sign in to your account
            </h1>

            <SigninForm />
        </div>
    );
};

export default page;
