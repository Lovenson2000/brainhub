"use client";

import React from "react";
import {
  useForm,
  FormProvider,
  useFormContext,
  Resolver,
} from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import Link from "next/link";
import { useRouter } from "next/navigation";

// Define validation schemas for each step
const Step1Schema = z.object({
  firstname: z.string().min(1, "First name is required"),
  lastname: z.string().min(1, "Last name is required"),
  email: z.string().email("Please enter a valid email address"),
});

const Step2Schema = z
  .object({
    password: z.string().min(6, "Password must be at least 6 characters"),
    confirmPassword: z
      .string()
      .min(6, "Confirm Password must be at least 6 characters"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

const Step3Schema = z.object({
  school: z.string().optional(),
  major: z.string().optional(),
  bio: z.string().optional(),
});

// Define a type for the resolver context
type ResolverContext = { step: number; submitter: string };

// Resolver function to validate based on the current step
const resolver: Resolver<any> = async (data, context) => {
  const ctx: ResolverContext = (context).current;
  let errors = {};

  if (ctx.step === 2 && (!data.password || !data.confirmPassword)) {
    errors = { ...errors, password: "Password is required" };
  }

  if (ctx.step === 3 && (!data.school || !data.major || !data.bio)) {
    errors = { ...errors, school: "School is required" };
  }

  return { values: data, errors };
};

export default function SignupForm() {
  const [step, setStep] = React.useState(1);
  const context = React.useRef({ step, submitter: "" });
  const methods = useForm({
    resolver,
    mode: "all",
    context,
  });
  const router = useRouter();

  const { register, handleSubmit } = methods;

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const submitter = context.current.submitter;
    const doHandleSubmit = handleSubmit(
      // onValid callback
      (data) => {
        const newStep = submitter === "prev" ? step - 1 : step + 1;

        if (newStep > 3) {
          // Final submit
          console.log("Final Data Before Submission:", data);
          // Make the POST request here
          fetch("http://localhost:8080/api/register", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
          })
            .then((response) => {
              if (response.ok) {
                toast({
                  title: "Success",
                  description: "Your registration was successful!",
                });
                router.push("/login");
              } else {
                return response.json().then((error) => {
                  throw new Error(error.error || "An error occurred");
                });
              }
            })
            .catch((error) => {
              toast({
                title: "Error",
                description:
                  error.message ||
                  "An error occurred while submitting the form.",
              });
            });
        } else {
          setStep(newStep);
          context.current.step = newStep;
        }
      },

      // onInvalid callback
      () => {}
    );
    doHandleSubmit(e);
  };

  const handlePrevStep = () => {
    context.current.submitter = "prev";
    onSubmit(new Event("submit"));
  };

  const handleNextStep = () => {
    context.current.submitter = "next";
    onSubmit(new Event("submit"));
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <FormProvider {...methods}>
        <form
          onSubmit={onSubmit}
          className="w-[20rem] md:w-[26rem] space-y-6 bg-white p-8 rounded-lg shadow-md"
        >
          <h2 className="text-2xl text-slate-800 font-bold text-center">
            Create Your Account
          </h2>

          <div className="relative flex justify-center mb-4">
            <div className="absolute right-0 text-gray-500">
              <span>{step} / 3</span>
            </div>
          </div>

          {step === 1 && (
            <>
              <InputField
                name="firstname"
                label="First Name"
                placeholder="First Name"
                register={register}
              />
              <InputField
                name="lastname"
                label="Last Name"
                placeholder="Last Name"
                register={register}
              />
              <InputField
                name="email"
                label="Email"
                placeholder="Email"
                type="email"
                register={register}
              />
            </>
          )}

          {step === 2 && (
            <>
              <InputField
                name="password"
                label="Password"
                placeholder="Password"
                type="password"
                register={register}
              />
              <InputField
                name="confirmPassword"
                label="Confirm Password"
                placeholder="Confirm Password"
                type="password"
                register={register}
              />
            </>
          )}

          {step === 3 && (
            <>
              <InputField
                name="school"
                label="School"
                placeholder="Harvard University"
                register={register}
              />
              <InputField
                name="major"
                label="Major"
                placeholder="Computer Science"
                register={register}
              />
              <InputField
                name="bio"
                label="Bio"
                placeholder="Building cool technologies..."
                register={register}
              />
            </>
          )}

          <div className="flex justify-between mt-4">
            {step > 1 && (
              <Button type="button" onClick={handlePrevStep}>
                Back
              </Button>
            )}
            {step < 3 ? (
              <Button type="button" onClick={handleNextStep}>
                Next
              </Button>
            ) : (
              <Button type="submit">Submit</Button>
            )}
          </div>

          <div className="mt-4 text-center">
            <h2 className="text-slate-800">Already have an account?</h2>
            <Link href="login" className="text-dark-blue">
              Log in
            </Link>
          </div>
        </form>
      </FormProvider>
    </div>
  );
}

// Component for rendering individual input fields
interface InputFieldProps {
  name: string;
  label: string;
  placeholder: string;
  type?: string;
  register: any; // Adjust the type as needed
}

function InputField({
  name,
  label,
  placeholder,
  type = "text",
  register,
}: InputFieldProps) {
  const {
    formState: { errors },
  } = useFormContext();
  const error = errors[name] as any;

  return (
    <div className="mb-4 relative">
      <label htmlFor={name} className="block text-slate-800">
        {label}
      </label>
      <input
        {...register(name)}
        id={name}
        placeholder={placeholder}
        type={type}
        className="w-full px-3 py-2 mt-1 border rounded-md"
      />
      <div className="absolute top-12 left-0 w-full text-red-500">
        <p
          className={`transition-opacity duration-300 ${
            error ? "opacity-100" : "opacity-0"
          }`}
        >
          {error?.message}
        </p>
      </div>
    </div>
  );
}
