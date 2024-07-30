"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "../components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../components/ui/form";
import { Input } from "../components/ui/input";
import { toast } from "../components/ui/use-toast";
import Link from "next/link";
import { useRouter } from "next/navigation";

const FormSchema = z.object({
  email: z.string().email({
    message: "Please enter a valid email address.",
  }),
  password: z.string().min(6, {
    message: "Password must be at least 6 characters.",
  }),
});

const API_BASE_URL = process.env.NGROK_URL || "http://localhost:8080";

export default function LoginForm() {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const router = useRouter();

  async function onSubmit(data: z.infer<typeof FormSchema>) {
    try {
      // Make the POST request to login
      const response = await fetch(`${API_BASE_URL}/api/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        const { token, user } = await response.json();
        localStorage.setItem("token", token)
        localStorage.setItem("user", JSON.stringify(user))


        toast({
          title: "Login successful!",
          description: "Redirecting you to the dashboard...",
        });

        router.push("/dashboard");
      } else {
        const error = await response.json();
        toast({
          title: "Login failed",
          description: error.message || "An error occurred during login.",
        });
      }
    } catch (error) {
      toast({
        title: "Error",
        description: "An unexpected error occurred.",
      });
    }
  }

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="w-[20rem] md:w-[26rem] space-y-6 bg-white p-8 rounded-lg shadow-md"
        >
          <h2 className="text-2xl text-slate-800 font-bold text-center">
            Log in to your account
          </h2>
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input type="email" placeholder="Email" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input type="password" placeholder="Password" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="">
            <h2 className="text-slate-800">Don't have an account yet?</h2>
            <Link href="signup" className="text-dark-blue">
              Sign up
            </Link>
          </div>
          <Button
            type="submit"
            className="w-full bg-light-blue  text-white py-2 rounded"
          >
            Log in
          </Button>
        </form>
      </Form>
    </div>
  );
}
