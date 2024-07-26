import Link from "next/link";
import { SiGoogleclassroom } from "react-icons/si";
import { GrAnalytics } from "react-icons/gr";
import { HiOutlineDocument } from "react-icons/hi";
import { GoTasklist } from "react-icons/go";
import { Button } from "./ui/button";

type SidebarItemProps = {
  label: string;
  icon: React.ReactNode;
  url: string;
};

export default function Sidebar() {
  return (
    <div className="w-64 relative bg-slate-50 border-r h-screen p-4">
      <div className=" p-4 bg-white shadow-md rounded-md">
        <h2 className="text-xl text-slate-900 text-blue-900">Brainhub</h2>
      </div>
      <nav className="mt-4">
        <SidebarItem
          label="Analytics"
          icon={<GrAnalytics size={16} />}
          url="/analytics"
        />
        <SidebarItem
          label="Tasks"
          icon={<GoTasklist size={20} />}
          url="/tasks"
        />
        <SidebarItem
          label="Sessions"
          icon={<SiGoogleclassroom size={18} />}
          url="/sessions"
        />
        <SidebarItem
          label="Documents"
          icon={<HiOutlineDocument size={20} />}
          url="/documents"
        />
      </nav>

      <Button
        type="submit"
        className="w-56 absolute bottom-4 bg-light-blue text-white py-2 rounded"
      >
        Go To Study Room
      </Button>
    </div>
  );
}

function SidebarItem({ label, icon, url }: SidebarItemProps) {
  return (
    <Link
      href={url}
      className="flex items-center px-4 py-2 hover:bg-white hover:border rounded-md transition-colors duration-200"
    >
      <span className="mr-4 text-slate-700 font-semibold">{icon}</span>
      <span className="text-slate-700 font-light hover:text-light-blue">
        {label}
      </span>
    </Link>
  );
}
