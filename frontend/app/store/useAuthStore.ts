import { create } from "zustand";

type AuthState = {
    accessToken: string | null;
    isLoggedIn: boolean;
    setAcessToken: (accessToken: string) => void;
    setIsLoggedIn: (isLoggedIn: boolean) => void;
    clearAccessToken: () => void;
    userID: string | undefined;
    setUserID: (id: string) => void;
};

export const useAuthStore = create<AuthState>((set) => ({
    accessToken: null,
    isLoggedIn: false,
    setAcessToken: (accessToken: string) => set({ accessToken: accessToken }),
    setIsLoggedIn: (isLoggedIn: boolean) => set({ isLoggedIn: isLoggedIn }),
    clearAccessToken: () =>
        set({ accessToken: null, isLoggedIn: false, userID: undefined }),
    userID: undefined,
    setUserID: (id: string) => set({ userID: id }),
}));
