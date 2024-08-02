import React from "react";

type TaskDescriptionProps = {
  description: string;
};
export default function TaskDescription({
  description,
}: TaskDescriptionProps) {
  return (
    <div>
      <p className="text-slate-800 m-1 font-light">{description}</p>
    </div>
  );
}
