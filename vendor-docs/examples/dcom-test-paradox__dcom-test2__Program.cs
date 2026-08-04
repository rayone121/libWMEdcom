using System;
using System.Collections.Generic;
using System.Dynamic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace dcom_test2
{
    class Program
    {
        static void Main(string[] args)
        {
            DocImpServer.DocImpObject o = new DocImpServer.DocImpObject();

            //    o.LogOn("GABI", "1");
            if (args.Length == 0)
                return;

            string name = args[0];

            if (name == "GetListaFirme")
            {
                string s = string.Join("\n", o.GetListaFirme());

                Console.Write(s);
            }

            if (name == "GetListaLuni")
            {
                if (args.Length < 2)
                    Console.Write("parametri insuficienti pentru GetListaLuni\n");

                string firma = args[1];

                o.SetNumeFirma(firma);

                dynamic arr = o.GetListaLuni(firma);
                if (arr == null)
                {
                    Console.Write("Nu am gasit lista luni pentru " + firma + "\n");
                }
                else
                {
                    string s = string.Join("\n", o.GetListaLuni(firma));
               //   Console.Write("Lista luni pentru " + firma + " este:\n");
                    Console.Write(s);
                }

            }

            if (name == "GetListaGestiuni")
            {
                if (args.Length < 2)
                {
                    Console.Write("parametri insuficienti pentru GetListaGestiuni\n");
                    return;
                }

                string firma = args[1];

                o.SetNumeFirma(firma);

                int err;
                dynamic arr = o.GetListaGestiuni(out err);
                //       dynamic arr2 = o.GetListaCatPret(out x);

                if (arr == null)
                {
                    Console.Write("Nu am gasit lista gestiuni pentru " + firma + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            if (name == "GetStocArticol")
            {
                if (args.Length < 4)
                {
                    Console.Write("parametri insuficienti pentru GetStocArticol\n");
                    return;
                }

                string firma = args[1];
                o.SetNumeFirma(firma);

                string gestID = args[2];
                string artID = args[3];

                int err;
                dynamic arr = o.GetStocArticol(artID, gestID, out err);
                if (arr == null)
                {
                    Console.Write("Nu am gasit stocul pt articolul " + artID + " si gestiunea " + gestID + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            if (name == "GetPretVanzare")
            {
                if (args.Length < 4)
                {
                    Console.Write("parametri insuficienti pentru GetStocArticol\n");
                    return;
                }

                string firma = args[1];
                o.SetNumeFirma(firma);

                string artID = args[2];
                string partID = args[3];

                int err;
                dynamic arr = o.GetPretVanzare(artID, partID, out err);
                if (arr == null)
                {
                    Console.Write("Nu am gasit pretul pt articolul " + artID + " si partenerul " + partID + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            if (name == "AdaugaPartener")
            {
                if (args.Length < 3)
                {
                    Console.Write("parametri insuficienti pentru AdaugaPartener\n");
                    return;
                }
                string firma = args[1];
                o.SetNumeFirma(firma);

                //         string infoPart = "106;SC MAREX SRL;RO544344;Iasi;str culturii nr 71 bis;232712844;Ionescu Adrian;;;1;J212294;;BRD;;;111;;;;;;;105;;;;;;;;";
                string filePath = args[2];
                string infoPart = System.IO.File.ReadAllText(filePath);
                int res = o.AdaugaPartener(infoPart);

                if (res == 0)
                    Console.Write("Partenerul nu a putut fi adaugat\n");
                else
                {
                    Console.Write("Partenerul a fost adaugat\n");
                    int err;
                    dynamic arr = o.GetListaParteneri(out err);

                    if (arr == null)
                    {
                        Console.Write("Nu am gasit lista parteneri pentru " + firma + "\n");
                    }
                    else
                    {
                        string s = string.Join("\n", arr);
                        Console.Write(s);
                    }
                }
            }

            if (name == "GetListaParteneri")
            {
                if (args.Length < 2)
                {
                    Console.Write("parametri insuficienti pentru GetListaParteneri\n");
                    return;
                }
                string firma = args[1];
                o.SetNumeFirma(firma);

                int err;
                dynamic arr = o.GetListaParteneri(out err);

                if (arr == null)
                {
                    Console.Write("Nu am gasit lista parteneri pentru " + firma + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            //GetStocArticole firma
            if (name == "GetStocArticole")
            {
                if (args.Length < 2)
                {
                    Console.Write("parametri insuficienti pentru GetStocArticole\n");
                    return;
                }
                string firma = args[1];
                o.SetNumeFirma(firma);

                int err;
                dynamic arr = o.GetStocArticole(out err);

                if (arr == null)
                {
                    Console.Write("Nu am gasit stoc articole pentru " + firma + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            //ImportaComenzi firma caleFisier
            if (name == "ImportaComenzi")
            {
                if (args.Length < 3)
                {
                    Console.Write("parametri insuficienti pentru ImportaComenzi\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

                //citire
                string filePath = args[2];
                int res = o.GetDocFromFile(filePath);
                if (res == 0)
                {
                    Console.Write("Eroare la citirea fisierului " + filePath + "\n");
                    return;
                }
                Console.Write("citire ok\n");

                //validare
                res = o.ComenziValide();
                dynamic le = o.GetListaErori();
                if (res == 0)
                {
                    Console.Write("Eroare la validare pt fisierul " + filePath + "\n");
                    string les = string.Join("\n", le);
                    Console.Write(les);
                    return;
                }
                Console.Write("validare ok\n");

                //import
                res = o.ImportaComenzi();
                if (res == 0)
                {
                    Console.Write("Eroare la import comenzi din fisierul " + filePath + "\n");
                    return;
                }
                Console.Write("import ok\n");
            }


            //ImportaFacturi firma caleFisier
            if (name == "ImportaFacturi")
            {
                if (args.Length < 3)
                {
                    Console.Write("parametri insuficienti pentru ImportaFacturi\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

                //citire
                string filePath = args[2];
                int res = o.GetDocFromFile(filePath);
                if (res == 0)
                {
                    Console.Write("Eroare la citirea fisierului " + filePath + "\n");
                    return;
                }
                Console.Write("citire ok\n");

                //validare
                res = o.DateValide();
                dynamic le = o.GetListaErori();
                if (res == 0)
                {
                    Console.Write("Eroare la validare pt fisierul " + filePath + "\n");
                    string les = string.Join("\n", le);
                    Console.Write(les);
                    return;
                }
                Console.Write("validare ok\n");

                //import
                res = o.ImportaFacturi();
                if (res == 0)
                {
                    Console.Write("Eroare la import facturi din fisierul " + filePath + "\n");
                    return;
                }
                Console.Write("import ok\n");
            }


            //1. AddProduct firma filePath
            if (name == "AddProduct")
            {
                if (args.Length < 3)
                {
                    Console.Write("parametri insuficienti pentru AddProduct\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

                //file
                string infoProd = System.IO.File.ReadAllText(args[2]);

                /*
                int res = o.AddProduct(infoProd);
                if (res == 0)
                    Console.Write("Produsul nu a putut fi adaugat\n");
                else
                {
                    Console.Write("Produsul a fost adaugat\n");
                }
                */
            }



            // 2. ImportaBonuriConsum firma filePath 
            // 3. BonuriConsumValide
            // ImportaBonuriConsum bbbb e:/bonuri.txt
            if (name == "ImportaBonuriConsum")
            {
                if (args.Length < 3)
                {
                    Console.Write("parametri insuficienti pentru ImportaBonuriConsum\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

                //luna lucru
                o.SetLunaLucru(2016, 10);

                //citire
                string filePath = args[2];
                int res = o.GetDocFromFile(filePath);
                if (res == 0)
                {
                    Console.Write("Eroare la citirea fisierului " + filePath + "\n");
                    return;
                }
                Console.Write("citire ok\n");

                //validare
                res = o.BonuriConsumValide();
                dynamic le = o.GetListaErori();
                if (res == 0)
                {
                    Console.Write("Eroare la validare bonuri pt fisierul " + filePath + "\n");
                    string les = string.Join("\n", le);
                    Console.Write(les);
                    return;
                }
                Console.Write("validare bonuri ok\n");

                //import
                res = o.ImportaBonuriConsum();
                if (res == 0)
                {
                    Console.Write("Eroare la ImportBonuriConsum din fisierul " + filePath + "\n");
                    return;
                }
                Console.Write("import ok\n");
            }

            // 4. GetStocuriPeGestiuni firma ok
            // GetStocuriPeGestiuni bbbb
            if (name == "GetStocuriPeGestiuni")
            {
                if (args.Length < 2)
                {
                    Console.Write("parametri insuficienti pentru GetStocuriPeGestiune\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

                int err;
                dynamic arr = o.GetStocuriPeGestiuni(out err);

                if (arr == null)
                {
                    Console.Write("Nu am gasit stocuri pe gestiune pentru " + firma + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            //5. GetTransferuri firma dataStart dataEnd
            // GetTransferuri bbbb 10-OCT-2012 10-OCT-2016
            if (name == "GetTransferuri")
            {
                if (args.Length < 2)
                {
                    Console.Write("parametri insuficienti pentru GetTransferuri\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

          //      string dataStart = args[2];
        //        string dataEnd = args[3];

                int err;
                dynamic arr = o.GetTransferuri(out err);
                if (arr == null)
                {
                    Console.Write("Nu am gasit transferuri pentru " + firma + "\n");
                }
                else
                {
                    string s = string.Join("\n", arr);
                    Console.Write(s);
                }
            }

            //6. ImportaTransferuri firma filePath
            // ImportaTransferuri bbbb "e:/transf"
            if (name == "ImportaTransferuri")
            {
                if (args.Length < 3)
                {
                    Console.Write("parametri insuficienti pentru ImportaTransferuri\n");
                    return;
                }

                //firma de lucru
                string firma = args[1];
                o.SetNumeFirma(firma);

                //luna lucru
                o.SetLunaLucru(2016, 10);

                //citire
                string filePath = args[2];
                int res = o.GetDocFromFile(filePath);
                if (res == 0)
                {
                    Console.Write("Eroare la citirea fisierului " + filePath + "\n");
                    return;
                }
                Console.Write("citire ok\n");

                res = o.TransferuriValide();
                dynamic le = o.GetListaErori();
                if (res == 0)
                {
                    Console.Write("Eroare la validare pt fisierul " + filePath + "\n");
                    string les = string.Join("\n", le);
                    Console.Write(les);
                    return;
                }
                Console.Write("validare ok\n");

                //import
                res = o.ImportaTransferuri();
                if (res == 0)
                {
                    Console.Write("Eroare la ImportaTransferuri din fisierul " + filePath + "\n");
                    return;
                }
                Console.Write("import ok\n");
            }


            if (name == "ActualizeazaAcceptat")
            {
                string firma = args[1];
                o.SetNumeFirma(firma);
                string an1 = args[2];
                string luna1 = args[3];
                o.SetLunaLucru(int.Parse(an1), int.Parse(luna1));

                string codcomanda = args[4];
                string datacomanda = args[5];
                string idarticol = args[6];
                string cantitate = args[7];

                //     string obj = codcomanda + ';' + datacomanda + ';' + idarticol + ';' + cantitate;
                //     int res = o.ActualizeazaAcceptat(obj);

                /*
                dynamic obj = new ExpandoObject();
                obj.codcomanda = args[4];
                obj.datacomanda = args[5];
                obj.idarticol = args[6];
                obj.cantitate = args[7];
                */

                string[] obj = new string[] { codcomanda, datacomanda, idarticol, cantitate};


                int res = o.ActualizeazaAcceptat(obj);

                if (res == 0)
                {
                    Console.Write("Articolul nu a putut fi modificat\n");
                }
                else
                {
                    Console.Write("Comanda a fost modificata\n");
                }
            }




        }

    }
}
